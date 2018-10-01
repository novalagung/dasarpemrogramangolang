package main

import "fmt"
import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func connect() (*mgo.Session, error) {
	var session, err = mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return session, nil
}

func remove() {
	var session, err = connect()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	defer session.Close()
	var collection = session.DB("belajar_golang").C("student")

	var selector = bson.M{"name": "John Wick"}
	err = collection.Remove(selector)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println("Remove success!")
}

func main() {
	remove()
}
