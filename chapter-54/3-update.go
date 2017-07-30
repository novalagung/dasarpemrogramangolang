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

func update() {
	var session, err = connect()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	defer session.Close()
	var collection = session.DB("belajar_golang").C("student")

	var selector = bson.M{"name": "Wick"}
	var changes = student{"John Wick", 2}
	err = collection.Update(selector, changes)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println("Update success!")
}

func main() {
	update()
}
