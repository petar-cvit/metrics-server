package main

import (
	"example.com/assignment/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	metrics := internal.NewMetricsStorage()

	r.GET("/metrics", func(c *gin.Context) {
		c.String(http.StatusOK, metrics.GetMetrics())
	})

	r.Run(":12345")
}
