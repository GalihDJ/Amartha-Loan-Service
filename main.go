package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	docs "amartha-loan-service/docs"
	conn "amartha-loan-service/utils/Connections"

	loan "amartha-loan-service/api/v1/Loan"
)

// @title           Amartha Loan Service
// @version			0.0.1
// @description     Swagger documentation for Amartha Loan Service API

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

type Object struct {
	Str string
	Num int
}

func main() {
	fmt.Println(fmt.Sprintf("Welcome to Amartha Loan Service!"))

	router := gin.Default()

	// swagger basepath and host
	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASE_PATH")
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	fmt.Println("Host: ", docs.SwaggerInfo.Host)

	// connect PSQL
	connPSQL := conn.ConnectionMapPSQL[os.Getenv("LOAN_SERVICE_PLATFORM_ENVIRONMENT")]
	fmt.Println("PSQL DB: " + connPSQL.Database)

	// initialize router for modules
	loan.InitializeLoan(router, &connPSQL)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", "3000"), router); err != nil {
		log.Fatal(err)
	}
}
