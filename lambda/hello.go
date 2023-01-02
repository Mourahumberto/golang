package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest() (string, error) {
	return "Humberto Moura!", nil
}
func main() {
	lambda.Start(handleRequest)
}
