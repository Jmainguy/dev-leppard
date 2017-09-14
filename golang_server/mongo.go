package main

import (
    "gopkg.in/mgo.v2"
)

func insertMessage (mongoSession *mgo.Session, mongo_db string, values map[string]interface {}) {
    sessionCopy := mongoSession.Copy()
    coll := sessionCopy.DB(mongo_db).C("numbers")
    err := coll.Insert(values)
    check(err)
    sessionCopy.Close()
}

func mapPageToNumber(mongoSession *mgo.Session, mongo_db string, phoneNumber, pageId, numberId string) {
    sessionCopy := mongoSession.Copy()
    coll := sessionCopy.DB(mongo_db).C("pagemap")

    // Create map, and assign values
    values := make(map[string]interface{})
    values["phoneNumber"] = phoneNumber
    values["_id"] = pageId
    values["numberId"] = numberId

    // Insert into mongo
    _, err := coll.UpsertId(&pageId, values)
    check(err)
    sessionCopy.Close()
}

