package resp

import "github.com/aws/aws-lambda-go/events"

func OKResponse(body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: 200,
	}

}
