package main

import (
    "time"
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


type NumberJSON []struct {
    Number         string `json:"number"`
    City           string `json:"city"`
    State          string `json:"state"`
    NationalNumber string `json:"nationalNumber"`
    Price          string `json:"price"`
    Location       string `json:"location"`
}
