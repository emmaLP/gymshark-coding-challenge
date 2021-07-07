package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

type Data struct {
	ItemsOrdered int   `json:"items"`
	Packs        []int `json:"packSizes"`
}

type Response struct {
	Packs map[int]int
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)
	data := &Data{}
	headers := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		"Access-Control-Allow-Methods": "GET, POST, OPTIONS",
	}

	_ = json.Unmarshal([]byte(request.Body), data)
	if data.ItemsOrdered < 1 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "ItemsOrdered must be set or more than 0",
			Headers:    headers,
		}, nil
	} else if data.Packs == nil || len(data.Packs) == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "No packs available!",
			Headers:    headers,
		}, nil
	}

	var jsonData = determinePacks(data)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jsonData),
		Headers:    headers,
	}, nil
}

func main() {
	lambda.Start(Handler)

}
