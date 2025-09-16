package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.S3Event) (finished bool, err error) {

	fmt.Printf("----- start -----\n")
	println("Hello World From lambda Convert From Elbon!")
	return true, nil
}

func main() {
	lambda.Start(handler)
}
