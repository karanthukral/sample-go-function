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

	connStr := os.Getenv("DB_URL")
	fmt.Println("Connecting to db...")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("pinging...")
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("ponging...")

	msg["body"] = fmt.Sprintf("%s. Pinged db", msg["body"])

	return msg
}
