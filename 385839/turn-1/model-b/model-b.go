import (
	"fmt"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Create a custom collector
type reqHandler struct {
	reqCount prometheus.Counter
	reqLat   prometheus.Histogram
}

func newReqHandler() *reqHandler {
	return &reqHandler{
		reqCount: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "file_storage_request_count",
			Help: "Number of handled requests",
		}),
		reqLat: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name: "file_storage_request_duration_seconds",
			Help: "Duration of requests in seconds",
			Buckets: []float64{
				0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10,
			},
		}),
	}
}

func (rh *reqHandler) handleFileRequest(w http.ResponseWriter, r *http.Request) {
	// start time
	start := time.Now()

	// Your file handling logic here
	// ...

	// Measure request duration
	duration := time.Since(start).Seconds()

	// Increment the request counter and observe latency
	rh.reqCount.Inc()
	rh.reqLat.Observe(duration)
}

func main() {
	rh := newReqHandler()

	prometheus.MustRegister(rh.reqCount)
	prometheus.MustRegister(rh.reqLat)

	// Serve metrics at /metrics
	http.HandleFunc("/metrics", promhttp.Handler.ServeHTTP)

	// Handle other routes
	http.HandleFunc("/handle-file", rh.handleFileRequest)

	fmt.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
