package main

import (
	"fixerapi/fixer"
	"fixerapi/middlewares"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	_ = godotenv.Load()
	_ = sentry.Init(sentry.ClientOptions{Dsn: os.Getenv("SENTRY")})
	fixer.SetKey(os.Getenv("FIXER_KEY"))
	fixer.Update()
	middlewares.SetKey()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middlewares.ApiKeyMiddleware())
	fixer.Routes(r)
	r.Run(":8000")
}
