package main

import (

    "net/http"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strings"
    "bytes"

)

func orderNumber (token, secret string) (number, numberId string) {

    // Have to choose an area, North Carolina seems like a good one.
    // We only will ever be ordering one number at a time
    url := "https://api.catapult.inetwork.com/v1/availableNumbers/local?state=NC&quantity=1"

    req, err := http.NewRequest("POST", url, nil)
    check(err)

    req.SetBasicAuth(token, secret)
    resp, err := http.DefaultClient.Do(req)
    check(err)
    defer resp.Body.Close()
    numberJson := NumberJSON{}
    err = json.NewDecoder(resp.Body).Decode(&numberJson)
    check(err)
    fmt.Println(numberJson[0].Number)
    locationArray := strings.Split(numberJson[0].Location, "/")
    numberId = locationArray[len(locationArray) -1]
    fmt.Println(numberId)
    number = numberJson[0].Number

    return
}

type ApplicationId struct{
    applicationId string
}

func assignNumber (user, token, secret, number, appId, numberId string) {

    url := fmt.Sprintf("https://api.catapult.inetwork.com/v1/users/%s/phoneNumbers/%s", user, number)
    //url := "https://requestb.in/17jsu5d1"
    fmt.Println(url)

    b := fmt.Sprintf(`{"applicationId": "%s"}`, appId)
    var jsonStr = []byte(b)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    check(err)

    req.SetBasicAuth(token, secret)
    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    check(err)
    defer resp.Body.Close()
    bb, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(bb))
    fmt.Println(resp.StatusCode)
    
}
