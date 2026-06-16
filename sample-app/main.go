package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "app_request_duration_seconds",
			Help:    "HTTP request latency in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	processedItems = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "app_processed_items",
			Help: "Number of items processed",
		},
		[]string{"type"},
	)

	errorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_errors_total",
			Help: "Total number of errors",
		},
		[]string{"type"},
	)

	cpuLoad = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "app_cpu_load",
			Help: "Application CPU load simulation",
		},
	)
)

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(processedItems)
	prometheus.MustRegister(errorCounter)
	prometheus.MustRegister(cpuLoad)
}

func recordMetrics() {
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			// Simulate metrics
			processedItems.WithLabelValues("type_a").Set(float64(rand.Intn(100)))
			processedItems.WithLabelValues("type_b").Set(float64(rand.Intn(100)))

			cpuLoad.Set(math.Sin(float64(time.Now().Unix()%60)) * 50 + 50)

			if rand.Float64() > 0.9 {
				errorCounter.WithLabelValues("type_error").Inc()
			}
		}
	}()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		requestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
	}()

	status := 200
	if rand.Float64() > 0.95 {
		status = 500
		errorCounter.WithLabelValues("request_error").Inc()
	}

	requestCounter.WithLabelValues(r.Method, r.URL.Path, fmt.Sprintf("%d", status)).Inc()

	w.WriteHeader(status)
	fmt.Fprintf(w, "Hello from sample app! Status: %d\n", status)
}

func main() {
	recordMetrics()

	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "healthy")
	})
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Starting sample app on :8888")
	http.ListenAndServe(":8888", nil)
}
