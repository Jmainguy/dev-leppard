package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    //"strings"
    "time"
    "net/http"
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
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.POST("/messages", func(c *gin.Context) {
        var json MessageJSON
        c.Bind(&json)

        bjson := json[0]
        //check the 'to' number in the json
        fmt.Println(bjson.To)

        //find the page ID associated with that 'to' number

        //check the media urls in the json
        fmt.Println(bjson.Message.Media)

        //update the content on that page with the media

        c.JSON(200, bjson)
    })
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
    server.Run(":25551")
}

