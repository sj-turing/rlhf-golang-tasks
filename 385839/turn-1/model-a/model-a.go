package main

/*

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_latency_seconds",
			Help:    "Request latency in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)

func init() {
	prometheus.MustRegister(requestLatency)
}

func main() {
	router := gin.Default()

	router.GET("/metrics", promhttp.Handler())

	router.GET("/files/:path", func(c *gin.Context) {
		method := c.Request.Method
		start := prometheus.Now()

		// Handle file request logic here

		latency := prometheus.Since(start)
		requestLatency.WithLabelValues(method).Observe(latency.Seconds())

		c.JSON(200, gin.H{"message": "File found"})
	})

	router.Run(":8080")
}
*/
