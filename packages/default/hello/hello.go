package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/xo/dburl"
)

func Main(args map[string]interface{}) map[string]interface{} {
	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!"

	// Open up our database connection.
	connection := os.Getenv("DB_URL")
	dbURL, err := dburl.Parse(connection)
	if err != nil {
		panic(err)
	}

	dbPassword, _ := dbURL.User.Password()
	dbName := strings.Trim(dbURL.Path, "/")
	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=true", dbURL.User.Username(), dbPassword, dbURL.Hostname(), dbURL.Port(), dbName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

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
