package middlewares

import (
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var apiKey string

func SetKey() {
	apiKey = os.Getenv("API_KEY")
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.AbortWithStatusJSON(code, resp)
}

func ApiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userKey, _ := c.GetQuery("apikey")
		if userKey == "" {
			respondWithError(http.StatusUnauthorized, "API token required", c)
			sentry.CaptureException(errors.New("API token required"))
			return
		}
		if userKey != apiKey {
			respondWithError(http.StatusUnauthorized, "Invalid API token", c)
			sentry.CaptureException(errors.New("Invalid API token"))
			return
		}
		c.Next()
	}
}
