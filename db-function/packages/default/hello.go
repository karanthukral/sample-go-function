package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Main(args map[string]interface{}) map[string]interface{} {
	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!"

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, err := db.Query("SHOW tables")
	if err != nil {
		panic(err.Error())
	}

	tableCount := 0

	for res.Next() {
		tableCount++
	}

	msg["body"] = fmt.Sprintf("%s\nTable Count: %d", msg["body"], tableCount)

	return msg
}
