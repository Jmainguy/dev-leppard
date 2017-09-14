package main

import (
    "io/ioutil"
    "github.com/ghodss/yaml"
)

type Config struct {
    User string `json:"user"`
    Token string `json:"token"`
    Secret string `json:"secret"`
    Mongo_db string `json:"mongo_db"`
    Mongo_addr string `json:"mongo_addr"`
    AppId string `json:"appId"`
}

func config() (user, token, secret, mongo_db, mongo_addr, appId string){
    var v Config
    config_file, err := ioutil.ReadFile("/opt/messentation/config.yaml")
    check(err)
    yaml.Unmarshal(config_file, &v)
    user = v.User
    token = v.Token
    secret = v.Secret
    mongo_db = v.Mongo_db
    mongo_addr = v.Mongo_addr
    appId = v.AppId
    return
}
