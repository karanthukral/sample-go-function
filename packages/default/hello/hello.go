package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Main(args map[string]interface{}) map[string]interface{} {
	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!"

	caCert := os.Getenv("CA_CERT")
	fmt.Println(caCert)
	if caCert != "" {
		msg["body"] = fmt.Sprintf("%s\n found CA_CERT", msg["body"])
	}

	opts := options.Client()
	opts.ApplyURI(os.Getenv("DB_URL"))

	roots := x509.NewCertPool()
	// ok = roots.AppendCertsFromPEM([]byte(caCert))
	roots.AppendCertsFromPEM([]byte(caCert))
	// if !ok {
	// 	panic("cert didn't work")
	// }
	opts.SetTLSConfig(&tls.Config{
		RootCAs: roots,
	})

	client, err := mongo.NewClient(opts)
	if err != nil {
		fmt.Printf("client failed: %s", err.Error())
		// panic(err.Error())
	}

	fmt.Println("connecting....")
	err = client.Connect(context.TODO())
	if err != nil {
		fmt.Printf("failed connect: %s", err.Error())
	}

	fmt.Println("pinging...")
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		fmt.Printf("failed ping: %s", err.Error())
	} else {
		fmt.Println("ping worked")
	}

	// } else {
	// 	ctx := context.Background()
	// 	err = client.Connect(ctx)
	// 	if err != nil {
	// 		fmt.Printf("errored connecting mongo: %s", err.Error())
	// 	} else {
	// 		if err := client.Ping(ctx, readpref.Primary()); err != nil {
	// 			fmt.Printf("errored pinging mongo: %s", err.Error())
	// 			// panic(err)
	// 		}
	// 	}
	// }

	msg["body"] = fmt.Sprintf("%s\n Mongo Pinged", msg["body"])

	return msg
}
