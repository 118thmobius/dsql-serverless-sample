package main

import (
	"context"
	"encoding/json"
	"github.com/118thmobius/dsql-serverless-sample/internal/adapter/apigateway/resp"
	"github.com/118thmobius/dsql-serverless-sample/internal/infrastucture/db"
	"github.com/118thmobius/dsql-serverless-sample/internal/usecase"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var ctx context.Context
var accountUseCase usecase.AccountUseCase

func handler(requestCtx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse request
	var transferRequest TransferRequest
	err := json.Unmarshal([]byte(request.Body), &transferRequest)
	if err != nil {
		return resp.InternalServerErrorResponse(), err
	}

	// Process request
	tx, err := accountUseCase.Transfer(ctx, transferRequest.FromId, transferRequest.ToId, transferRequest.Amount)
	if err != nil {
		return resp.InternalServerErrorResponse(), err
	}

	// Return result
	resultJson, err := json.Marshal(NewTransferResponse(tx, "Transferred successfully."))
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
