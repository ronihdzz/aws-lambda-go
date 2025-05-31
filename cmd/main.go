package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"github.com/ronihdz/simple-gin-lambda/internal/router"
)

func main() {

	basePath := os.Getenv("ROOT_PATH")
	if basePath == "" {
		basePath = "/"
	}

	r := router.New(basePath)

	if _, inLambda := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME"); inLambda {
		lambda.Start(ginadapter.NewV2(r).ProxyWithContext)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("[	// Handler global para 404BOOT] Local en :%s (basePath=%s)", port, basePath)
	if err := r.Run(":" + port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[FATAL] %v", err)
	}
}
