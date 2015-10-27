package main

import "fmt"
import "database/sql"
import "os"
import _ "github.com/go-sql-driver/mysql"

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() *sql.DB {
	var db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	return db
}

func sqlQueryRow() {
	var db = connect()
	defer db.Close()

	var result = student{}
	var id = "E001"
	var err = db.QueryRow("select name, grade from tb_student where id = ?", id).Scan(&result.name, &result.grade)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Printf("name: %s\ngrade: %d\n", result.name, result.grade)
}

func main() {
	sqlQueryRow()
}
