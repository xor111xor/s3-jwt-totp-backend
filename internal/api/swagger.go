package api

import (

	// swagger embed files
	// gin-swagger middleware

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xor111xor/s3-jwt-totp-backend/docs"
)

func Swagger(router *gin.Engine, schema, host, port, path_api string) {
	// Swagger
	docs.SwaggerInfo.Title = "Swagger of Storage API"
	docs.SwaggerInfo.Description = "REST API server of storage with auth"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = host + ":" + port
	docs.SwaggerInfo.BasePath = "/" + path_api
	docs.SwaggerInfo.Schemes = []string{schema}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
