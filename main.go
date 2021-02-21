package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func() string {
		return "Hello, world!"
	})
}
