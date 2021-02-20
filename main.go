package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func() {
		fmt.Println("Hello, world!")
	})
}
