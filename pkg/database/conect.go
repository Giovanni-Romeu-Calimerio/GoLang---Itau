package database

import (
	"database/sql"
	"fmt"
	"internal/pkg/entity"
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

func conect() {

	db, err := sql.Open("sqlite3", "C:\\Users\\heder\\go\\src\\github.com\\GoLangCurso-1\\pkg\\database\\dbGolang.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//exec(db, `insert into products (Name, Code, Price, CreatedAt, UpdatedAt) values ( "Mala", 1017, 100.7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`)
	rows, _ := db.Query("SELECT ID, Name, Code, Price, UpdatedAt FROM products ")
	for rows.Next() {
		var p entity.Produto
		rows.Scan(&p.ID, &p.Name, &p.Code, &p.Price, &p.UpdatedAt)
		fmt.Println(p)
	}

}
