package main

import (
	"encoding/json"
	"net/http"
	"time"

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
	//router := gin.Default()

	//router.GET("/metrics", promhttp.Handler())

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/file/:path", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		start := time.Now()

		// Handle file request logic here
		time.Sleep(100 * time.Millisecond)

		latency := time.Since(start)
		requestLatency.WithLabelValues(method).Observe(latency.Seconds())

		//c.JSON(200, gin.H{"message": "File found"})

		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")

		type response struct {
			Message string `json:"message"`
		}

		res := response{
			Message: "File Found",
		}

		bb, _ := json.Marshal(res)

		w.Write(bb)

	})

	/*
		router.GET("/files/:path", func(c *gin.Context) {
			method := c.Request.Method
			start := time.Now()

			// Handle file request logic here

			latency := time.Since(start)
			requestLatency.WithLabelValues(method).Observe(latency.Seconds())

			c.JSON(200, gin.H{"message": "File found"})
		})

		router.Run(":8080")
	*/

	http.ListenAndServe(":8080", nil)
}
