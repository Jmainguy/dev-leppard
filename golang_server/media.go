package main

import (

    "net/http"
    "strings"
    "fmt"
    "os"
    "io"

)

func saveMedia (token, secret, media string) (imageUrl string) {
    
    req, err := http.NewRequest("GET", media, nil)
    check(err)

    req.SetBasicAuth(token, secret)

    resp, err := http.DefaultClient.Do(req)
    check(err)

    defer resp.Body.Close()
    // https://api.catapult.inetwork.com/v1/users/u-bul466c646zxrotsrn4qi7a/media/7_animated_happy1-m-5vjtsen3tpctcm6rnzxcqty.gif
    mediaNameArray := strings.Split(media, "/")
    mediaName := mediaNameArray[len(mediaNameArray) -1]
    savedMediaName := fmt.Sprintf("/opt/messentation/images/%s", mediaName)
    out, err := os.Create(savedMediaName)
    check(err)

    defer out.Close()
    io.Copy(out, resp.Body)

    imageUrl = fmt.Sprintf("/images/%s", mediaName)
    return
}
