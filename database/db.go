package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func DBInit(user string, password string,
	dbhost string, dbname string) {
	connectionString :=
		fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
			user,
			password,
			dbhost,
			dbname)

	fmt.Println(connectionString)

	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	CreateTableDummy(DB)
}

func CreateTableDummy(db *sql.DB) {
	if _, err := db.Exec(DummyTableCreate); err != nil {
		log.Fatal(err)
	}
}
