package resp

import "github.com/aws/aws-lambda-go/events"

func InternalServerErrorResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       "message: internal server error",
		StatusCode: 500,
	}
}
