package main

import (
	"context"
	"encoding/json"
	"github.com/118thmobius/dsql-serverless-sample/internal/adapter/apigateway/resp"
	"github.com/118thmobius/dsql-serverless-sample/internal/shared"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var internalServerError = events.APIGatewayProxyResponse{
	Body:       "message: Internal Server Error",
	StatusCode: 500,
}

func handler(requestCtx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Return result
	resultJson, err := json.Marshal(VersionResponse{
		State:   "Ready",
		Version: shared.Version(),
	})
	if err != nil {
		return internalServerError, err
	}

	return resp.OKResponse(string(resultJson)), nil
}

func main() {
	lambda.Start(handler)
}

type VersionResponse struct {
	State   string `json:"state"`
	Version string `json:"version"`
}
