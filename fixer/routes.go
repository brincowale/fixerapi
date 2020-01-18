package fixer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(route *gin.Engine) {
	var caches = route.Group("/fixer")
	{
		caches.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, getCurrencies())
		})
		caches.GET("/update", func(c *gin.Context) {
			c.JSON(http.StatusOK, Update())
		})
	}
}
