package main

import (
	"github.com/Experiment-Code/realworld-aws-lambda-dynamodb-go/service"
	"github.com/Experiment-Code/realworld-aws-lambda-dynamodb-go/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handle(input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	user, _, err := service.GetCurrentUser(input.Headers["Authorization"])
	if err != nil {
		return util.NewUnauthorizedResponse()
	}

	err = service.DeleteArticle(input.PathParameters["slug"], user.Username)
	if err != nil {
		return util.NewErrorResponse(err)
	}

	return util.NewSuccessResponse(200, nil)
}

func main() {
	lambda.Start(Handle)
}
