package main

import (
    "github.com/gin-gonic/gin"
    "strings"
    "time"
    "net/http"
    "gopkg.in/mgo.v2"
)

func main() {
    // Read Config, load values
    user, token, secret, mongo_db, mongo_addr, appId := config()
    // Open mongodb connection
    mongoDBDialInfo := &mgo.DialInfo{
        Addrs:    []string{mongo_addr},
        Timeout:  60 * time.Second,
    }
    // Open and maintain pool of sessions
    mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
    check(err)
    // Start callbackServer in a go function, so it runs in sep process
    go callbackServer(mongoSession, token, secret, mongo_db)
    // Let guiServer run as its own process, so this daemon has something that runs forever and doesnt die.
    guiServer(mongoSession, token, secret, mongo_db, user, appId)
}

func callbackServer(mongoSession *mgo.Session, token, secret, mongo_db string) {
    // Start gin-gonic webserver
    r := gin.Default()
    r.POST("/messages", func(c *gin.Context) {
        var newjson MessageJSON
        c.Bind(&newjson)

        bjson := newjson[0]
        //find the page ID associated with that 'to' number
        // Query db for pageid based on from

        // Save media
        for url, _ := range bjson.Message.Media {

            media := bjson.Message.Media[url]
            if strings.HasSuffix(media, "txt") {
                continue
            } else if strings.HasSuffix(media, "smil") {
                continue
            } else {
                imageUrl := saveMedia(token, secret, bjson.Message.Media[url])
                // Mongodb stuff below
                values := make(map[string]interface{})
                values["time"] = &bjson.Time
                values["from"] = &bjson.Message.From
                values["owner"] = &bjson.Message.Owner
                values["text"] = &bjson.Message.Text
                values["url"] = imageUrl
                //values["pageId"] = pageId
                values["pageId"] = "Jon"
                insertMessage(mongoSession, mongo_db, values)
            }
        }
        //update the content on that page with the media

        c.JSON(200, bjson)
    })

    // Run on port 25550
    r.Run(":25550")
}

func guiServer(mongoSession *mgo.Session, token, secret, mongo_db, user, appId string) {
    server := gin.Default()
    server.Static("/images", "/opt/messentation/images/")
    server.StaticFile("/", "templates/gui.html")
    server.POST("/pages", func(c *gin.Context) {
        //phoneNumber, numberId := orderNumber(token, secret)
        phoneNumber := "+19104270915"
        numberId := "n-gcqjn32ctbr2k7lfuxky6va"
        assignNumber(user, token, secret, phoneNumber, appId, numberId)
        pageId := randomString(10)
        mapPageToNumber(mongoSession, mongo_db, phoneNumber, pageId, numberId)
        c.HTML(http.StatusOK, "templates/pages.tmpl", gin.H{
            "pageId": pageId,
            "phoneNumber": phoneNumber,
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

