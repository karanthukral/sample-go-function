package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xo/dburl"
	"os"
)

func Main(args map[string]interface{}) map[string]interface{} {
	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!" + os.Getenv("DATABASE_URL")

	// Open up our database connection.
	dbURL, err := dburl.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}

	dbPassword, _ := dbURL.User.Password()
	dbName := strings.Trim(dbURL.Path, "/")
	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=true", dbURL.User.Username(), dbPassword, dbURL.Hostname(), dbURL.Port(), dbName)

	db, err := sql.Open("mysql", connectionString)
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
