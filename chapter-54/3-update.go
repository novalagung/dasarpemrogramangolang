package main

import "fmt"
import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"
import "os"

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func connect() *mgo.Session {
	var session, err = mgo.Dial("localhost")
	if err != nil {
		os.Exit(0)
	}
	return session
}

func update() {
	var session = connect()
	defer session.Close()
	var collection = session.DB("belajar_golang").C("student")

	var selector = bson.M{"name": "Wick"}
	var changes = student{"John Wick", 2}
	var err = collection.Update(selector, changes)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	update()
}
