package api

import (
	"time"

	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/xor111xor/s3-jwt-totp-backend/internal/domain"
)

// enable metrics for length cache & metrics gin
func PromethusMetrics(router *gin.Engine, duration time.Duration, cache domain.Cache) {
	p := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)

	p.AddCustomGauge("cache_length", "Cache length ", []string{})
	p.AddCustomGauge("no_verify_user", "Amount of unregistered users in cache", []string{})
	ticker := time.NewTicker(duration * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			length, _ := cache.LengthCache()
			notVerified, _ := cache.LengthUnverifiedUsers()
			p.SetGaugeValue("cache_length", []string{}, length)
			p.SetGaugeValue("no_verify_user", []string{}, notVerified)
		}
	}()

	router.Use(p.Instrument())
}
