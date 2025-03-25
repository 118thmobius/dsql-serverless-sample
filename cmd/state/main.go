package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/118thmobius/dsql-serverless-sample/internal/adapter/apigateway/resp"
	"github.com/118thmobius/dsql-serverless-sample/internal/domain"
	"github.com/118thmobius/dsql-serverless-sample/internal/infrastucture/db"
	"github.com/118thmobius/dsql-serverless-sample/internal/usecase"
	"github.com/google/uuid"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var ctx context.Context
var accountUseCase usecase.AccountUseCase

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Check Params
	userId, ok := request.PathParameters["userId"]
	if !ok {
		return resp.UserIdIsNotFoundResponse(), fmt.Errorf("userId is not found")
	}
	if _, err := uuid.Parse(userId); err != nil {
		return resp.InvalidUUIDFormatResponse(), fmt.Errorf("invalid UUID format: %w", err)
	}
	// Get Account data
	account, err := accountUseCase.GetAccountByID(ctx, userId)
	if err != nil {
		return resp.UserNotFoundResponse(), err
	}
	// Return Response
	response := AccountResponse{
		Account: account,
	}
	resultJson, err := json.Marshal(response)
	if err != nil {
		return resp.InternalServerErrorResponse(), err
	}
	return resp.OKResponse(string(resultJson)), nil
}

func main() {
	ctx = context.Background()
	pool, txManager, err := db.NewDSQLTransactionManager(ctx)
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	accountUseCase = usecase.NewAccountUseCase(*txManager)

	lambda.Start(handler)
}

type AccountResponse struct {
	Account domain.StatefulAccount `json:"account"`
}
