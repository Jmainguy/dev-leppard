package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.POST("/messages", func(c *gin.Context) {
        //check the 'to' number in the json
        //find the page ID associated with that 'to' number
        //check the media urls in the json
        //update the content on that page with the media
        c.String(200, "OK")
    })
    r.Run(":25550")
}
