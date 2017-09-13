package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "strings"
    "time"
    "net/http"
    "gopkg.in/mgo.v2"
)

type MessageJSON []struct {
    Type        string    `json:"type"`
    Time        time.Time `json:"time"`
    Description string    `json:"description"`
    To          string    `json:"to"`
    Message     struct {
        ID            string    `json:"id"`
        Time          time.Time `json:"time"`
        To            []string  `json:"to"`
        From          string    `json:"from"`
        Text          string    `json:"text"`
        ApplicationID string    `json:"applicationId"`
        Media         []string  `json:"media"`
        Owner         string    `json:"owner"`
        Direction     string    `json:"direction"`
    } `json:"message"`
}

func main() {
    // Start callbackServer in a go function, so it runs in sep process
    go callbackServer()
    // Let guiServer run as its own process, so this daemon has something that runs forever and doesnt die.
    guiServer()
}

func callbackServer() {
    // Read Config, load values
    _, token, secret, mongo_db, mongo_addr := config()

    // Open mongodb connection
    mongoDBDialInfo := &mgo.DialInfo{
        Addrs:    []string{mongo_addr},
        Timeout:  60 * time.Second,
    }

    // Create a session which maintains a pool of socket connections
    // to our MongoDB.
    mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
    check(err)

    // Start gin-gonic webserver
    r := gin.Default()
    r.POST("/messages", func(c *gin.Context) {
        var newjson MessageJSON
        c.Bind(&newjson)

        bjson := newjson[0]
        //check the 'to' number in the json
        fmt.Println(bjson.To)

        //find the page ID associated with that 'to' number

        // Save media
        for url, _ := range bjson.Message.Media {

            media := bjson.Message.Media[url]
            if strings.HasSuffix(media, "txt") {
                continue
            } else if strings.HasSuffix(media, "smil") {
                continue
            } else {
                imageUrl := saveMedia(token, secret, bjson.Message.Media[url])
                fmt.Println(imageUrl)
                // Mongodb stuff below
                values := make(map[string]interface{})
                values["pageID"] = "Jonboy"
                values["new"] = true
                values["time"] = &bjson.Time
                values["from"] = &bjson.Message.From
                values["owner"] = &bjson.Message.Owner
                values["text"] = &bjson.Message.Text
                values["url"] = imageUrl
                sessionCopy := mongoSession.Copy()
                coll := sessionCopy.DB(mongo_db).C("numbers")
                err = coll.Insert(values)
                check(err)
                sessionCopy.Close()
            }
        }
        //update the content on that page with the media

        c.JSON(200, bjson)
    })

    // Run on port 25550
    r.Run(":25550")
}

func guiServer() {
    server := gin.Default()
    server.StaticFile("/", "templates/gui.html")
    server.POST("/pages", func(c *gin.Context) {
        c.HTML(http.StatusOK, "templates/pages.tmpl", gin.H{
            "pageId": "foo",//TODO code to generate ID
            "phoneNumber": "+12223334444",//TODO code to order a number
        })
    })
    server.GET("/pages/:pageId", func(c *gin.Context) {
        pageId := c.Param("pageId")
        c.HTML(http.StatusOK, "templates/page.tmpl", gin.H{
            "pageId": pageId,
        })
    })
    server.Run(":25551")
}

