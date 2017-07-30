package main

import "fmt"
import "gopkg.in/mgo.v2"

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

func insert() {
	var session, err = connect()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	defer session.Close()

	var collection = session.DB("belajar_golang").C("student")
	err = collection.Insert(&student{"Wick", 2}, &student{"Ethan", 2})
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println("Insert success!")
}

func main() {
	insert()
}
