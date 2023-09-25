package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	ERROR_PERCENTAGE  = 0
	DELAY_IN_MS       = 0
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "The total number of http calls",
	}, []string{"verb", "statusCode", "route"})
	httpResponseSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_requests_time_seconds",
		Help: "The total time of http calls in seconds",
	}, []string{"verb", "statusCode", "route"})
)

func init() {
	errorLevel, err := strconv.Atoi(os.Getenv("ERROR_PERCENTAGE"))
	if err != nil {
		fmt.Printf("defaulting to [%d] error percentage\n", ERROR_PERCENTAGE)
	}

	delayLevel, err := strconv.Atoi(os.Getenv("ERROR_PERCENTAGE"))
	if err != nil {
		fmt.Printf("defaulting to [%d]ms delay\n", DELAY_IN_MS)
	}

	DELAY_IN_MS = delayLevel
	ERROR_PERCENTAGE = errorLevel

}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	n := rand.Intn(100)
	if n < ERROR_PERCENTAGE {
		httpRequestsTotal.WithLabelValues("GET", "500", "hello-world").Inc()
		httpResponseSeconds.WithLabelValues("GET", "500", "hello-world").Observe(time.Since(start).Seconds())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - BOOM!"))

		fmt.Printf("%s - 500\n", time.Now())
	} else {
		httpRequestsTotal.WithLabelValues("GET", "200", "hello-world").Inc()
		httpResponseSeconds.WithLabelValues("GET", "200", "hello-world").Observe(time.Since(start).Seconds())

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
		fmt.Printf("%s - 200\n", time.Now())
	}
}

func main() {
	http.HandleFunc("/hello-world", helloWorld)
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
