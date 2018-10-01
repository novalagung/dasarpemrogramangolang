package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/kidstuff/mongostore"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
)

const SESSION_ID = "id"

func newMongoStore() *mongostore.MongoStore {
	mgoSession, err := mgo.Dial("localhost:27123")
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(0)
	}

	dbCollection := mgoSession.DB("learnwebgolang").C("session")
	maxAge := 86400 * 7
	ensureTTL := true
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store := mongostore.NewMongoStore(
		dbCollection,
		maxAge,
		ensureTTL,
		authKey,
		encryptionKey,
	)
	return store
}

func main() {
	store := newMongoStore()

	e := echo.New()

	e.Use(echo.WrapMiddleware(context.ClearHandler))

	e.GET("/set", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		session.Values["message1"] = "hello"
		session.Values["message2"] = "world"
		session.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	e.GET("/get", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)

		if len(session.Values) == 0 {
			return c.String(http.StatusOK, "empty result")
		}

		return c.String(http.StatusOK, fmt.Sprintf(
			"%s %s",
			session.Values["message1"],
			session.Values["message2"],
		))
	})

	e.GET("/delete", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		session.Options.MaxAge = -1
		session.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	e.Logger.Fatal(e.Start(":9000"))
}
