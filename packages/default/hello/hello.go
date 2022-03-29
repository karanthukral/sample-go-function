package main

import (
	"fmt"
	"os"
)

func Main(args map[string]interface{}) map[string]interface{} {
	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!" + os.Getenv("DB_URL")

	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DB_URL")))
	// if err != nil {
	// 	fmt.Printf("errored connecting to mongo: %s", err.Error())
	// 	// panic(err.Error())
	// } else {
	// 	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
	// 		fmt.Printf("errored pinging mongo: %s", err.Error())
	// 		// panic(err)
	// 	}
	// }

	msg["body"] = fmt.Sprintf("%s\n Mongo Pinged", msg["body"])

	return msg
}
