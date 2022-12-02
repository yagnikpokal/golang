package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    float64 `json:"id`
	Value string  `json:"value`
}

type Responce struct {
	Message string `json:"message"`
	ok      bool   `json:"ok"`
}

func Handler(request Request) (Responce, error) {
	return Responce{
		Message: fmt.Sprintf("Process request ID: %f", request.ID),
		ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
