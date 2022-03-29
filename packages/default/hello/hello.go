package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"regexp"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Main(args map[string]interface{}) map[string]interface{} {
	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!"

	caCert := os.Getenv("CA_CERT")
	re := regexp.MustCompile(`-----BEGIN CERTIFICATE----- `)
	str := re.ReplaceAllString(caCert, "")

	re = regexp.MustCompile(` -----END CERTIFICATE-----`)
	str = re.ReplaceAllString(str, "")

	re = regexp.MustCompile(` `)
	str = re.ReplaceAllString(str, "\n")

	str = fmt.Sprintf("-----BEGIN CERTIFICATE-----%s-----END CERTIFICATE-----", str)

	// re = regexp.MustCompile(`\n-----END CERTIFICATE-----`)
	// str = re.ReplaceAllString(str, "-----END CERTIFICATE-----")
	fmt.Println(str)

	// roots := x509.NewCertPool()
	// ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	// if !ok {
	// 	panic("failed to parse root certificate")
	// }
	// if caCert != "" {
	// 	msg["body"] = fmt.Sprintf("%s\n found CA_CERT", msg["body"])
	// }

	opts := options.Client()
	opts.ApplyURI(os.Getenv("DB_URL"))

	roots := x509.NewCertPool()
	ok = roots.AppendCertsFromPEM([]byte(str))
	if !ok {
		panic("failed to parse cert")
	}
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
	err = client.Ping(context.TODO(), nil)
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
