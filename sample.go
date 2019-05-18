package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"os"
)

type Sample struct {
	UserID string `dynamo:"UserID,hash"`
	Name   string `dynamo:"Name"`
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String(os.Getenv("DYNAMO_ENDPOINT")),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		panic(err)
	}

	db := dynamo.New(sess)

	db.Table("Samples").DeleteTable().Run()

	err = db.CreateTable("Samples", Sample{}).Run()
	if err != nil {
		panic(err)
	}

	table := db.Table("Samples")
	err = table.Put(&Sample{UserID: "1", Name: "Test1"}).Run()
	if err != nil {
		panic(err)
	}

	var sample Sample
	err = table.Get("UserID", "1").One(&sample)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", sample)
}
