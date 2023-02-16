package database

import (
	"awesomeapiserver/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

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
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&model.Dummy{})
}
