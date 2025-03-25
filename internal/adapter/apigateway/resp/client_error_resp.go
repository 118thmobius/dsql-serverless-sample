package resp

import "github.com/aws/aws-lambda-go/events"

func UserIdIsNotFoundResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       "message: userId is not found",
		StatusCode: 400,
	}
}

func InvalidUUIDFormatResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       "message: invalid UUID format",
		StatusCode: 400,
	}
}

func UserNotFoundResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       "message: invalid UUID format",
		StatusCode: 400,
	}
}
