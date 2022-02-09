package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Main(args map[string]interface{}) map[string]interface{} {
	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!"

	// Open up our database connection.
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PASSWORD"), "verify-full")

	db, err := sql.Open("postgres", connectionString)
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
