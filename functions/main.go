package main

import (
	"context"
	handler "test3/module/item/controller"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()
	r.GET("/employee/:id", handler.GetSpecificUser)
	r.GET("/employees", handler.GetUsers)
	r.POST("/employee/:id", handler.PostUsers)

	r.DELETE("/employee/:id", handler.DelUsers)
	r.PUT("/employee/:id", handler.UpdateUser)

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
