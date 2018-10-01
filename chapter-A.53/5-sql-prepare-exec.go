package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into tb_student values (?, ?, ?, ?)")
	stmt.Exec("G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("insert success!")
}

func main() {
	sqlExec()
}
