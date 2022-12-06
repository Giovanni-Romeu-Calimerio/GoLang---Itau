package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

type product struct {
	id    int
	name  string
	code  int
	price float64
	datac string
	datau string
}

func main() {

	db, err := sql.Open("sqlite3", "C:\\Users\\heder\\go\\src\\github.com\\GoLangCurso\\cmd\\dbGolang.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//exec(db, `insert into products (name, code, price, datac, datau) values ( "Case", 1016, 10.7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`)

	rows, _ := db.Query("SELECT id, name, code, price, datau FROM products ")
	for rows.Next() {
		var p product
		rows.Scan(&p.id, &p.name, &p.code, &p.price, &p.datau)
		fmt.Println(p)
	}

}
