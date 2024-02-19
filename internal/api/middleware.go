package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xor111xor/s3-jwt-totp-backend/internal/domain"
	"github.com/xor111xor/s3-jwt-totp-backend/internal/utils"
)

func jwtEncode(config *domain.CommonConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		cookie, err := ctx.Cookie("token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		sub, err := utils.ValidateToken(token, config.SysConfig.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user, err := config.Cache.Get(fmt.Sprint(sub))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Token expired, please sign in"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}

// func bodySizeMiddleware(c *gin.Context) {
// 	var w http.ResponseWriter = c.Writer
// 	c.Request.Body = http.MaxBytesReader(w, c.Request.Body, maxBodyBytes)

// 	c.Next()
// }

func bodySizeMiddleware(max_mb int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		maxBytes := 1024 * 1024 * max_mb // 5MB
		var w http.ResponseWriter = ctx.Writer
		ctx.Request.Body = http.MaxBytesReader(w, ctx.Request.Body, maxBytes)

		ctx.Next()
	}
}
