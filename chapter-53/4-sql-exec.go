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

func sqlExec() {
	var db = connect()
	defer db.Close()

	var err error

	_, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", "G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = db.Exec("update tb_student set age = ? where id = ?", 28, "G001")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = db.Exec("delete from tb_student where id = ?", "G001")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	sqlExec()
}
