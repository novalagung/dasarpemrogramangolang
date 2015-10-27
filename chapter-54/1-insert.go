package main

import "fmt"
import "gopkg.in/mgo.v2"
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

func insert() {
	var session = connect()
	defer session.Close()
	var collection = session.DB("belajar_golang").C("student")

	var err = collection.Insert(&student{"Wick", 2}, &student{"Ethan", 2})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	insert()
}
