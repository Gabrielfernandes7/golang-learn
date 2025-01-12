package main

import (
	"database/sql"
	"log"
	"github.com/go-sql-driver/mysql"
)

func main()  {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}

	log.Println(db)
}