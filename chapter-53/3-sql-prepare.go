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

func sqlPrepare() {
	var db = connect()
	defer db.Close()

	var statement, err = db.Prepare("select name, grade from tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	var result1 = student{}
	statement.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result1.name, result1.grade)

	var result2 = student{}
	statement.QueryRow("W001").Scan(&result2.name, &result2.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result2.name, result2.grade)

	var result3 = student{}
	statement.QueryRow("B001").Scan(&result3.name, &result3.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result3.name, result3.grade)
}

func main() {
	sqlPrepare()
}
