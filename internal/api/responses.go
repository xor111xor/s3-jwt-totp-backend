package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, response{message})
}

type oTPResponse struct {
	Mail   string `json:"mail,omitempty"`
	OtpUri string `json:"otp_uri"`
}

type userResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Mail      string    `json:"mail,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Verified  bool      `json:"verified"`
}

type filesResponse struct {
	UUID     uuid.UUID `json:"uuid,omitempty"`
	Name     string    `json:"name,omitempty"`
	Uploaded time.Time `json:"uploaded"`
	Size     int64     `json:"size"`
}
