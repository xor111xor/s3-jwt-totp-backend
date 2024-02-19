package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xor111xor/s3-jwt-totp-backend/internal/domain"
)

func RunApi(config *domain.CommonConfig) error {
	// disable debug mode
	gin.SetMode(gin.ReleaseMode)

	// create Gin router
	router := gin.Default()

	// setting CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		"http://localhost:" + config.SysConfig.ServicePort,
		config.SysConfig.ServiceSchema + "://" + config.SysConfig.ServiceIP + ":" +
			config.SysConfig.ServicePort,
	}
	corsConfig.AllowCredentials = true

	router.Use(cors.New(corsConfig))

	// enable Swagger
	Swagger(router, config.SysConfig.ServiceSchema, config.SysConfig.ServiceIP,
		config.SysConfig.ServicePort, config.SysConfig.ServicePathAPI)

	// enable metrics
	PromethusMetrics(router, config.SysConfig.MetricsScrapeSec, config.Cache)

	// enable UI
	startUI(router, config.SysConfig.ServicePathUI)

	// enable API
	startAPI(router, config)

	// start the server
	err := router.Run(":" + config.SysConfig.ServicePort)
	if err != nil {
		return err
	}
	return nil
}
