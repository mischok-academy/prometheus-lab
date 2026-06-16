package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	mu              sync.Mutex
	requestCount    int64
	errorCount      int64
	processedItemsA int64 = 50
	processedItemsB int64 = 30
)

func recordMetrics() {
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			mu.Lock()
			processedItemsA = int64(rand.Intn(100))
			processedItemsB = int64(rand.Intn(100))
			mu.Unlock()

			if rand.Float64() > 0.95 {
				mu.Lock()
				errorCount++
				mu.Unlock()
			}
		}
	}()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	status := 200
	if rand.Float64() > 0.95 {
		status = 500
		mu.Lock()
		errorCount++
		mu.Unlock()
	}

	mu.Lock()
	requestCount++
	mu.Unlock()

	w.WriteHeader(status)
	fmt.Fprintf(w, "Hello from sample app! Status: %d\n", status)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	cpuLoad := math.Sin(float64(time.Now().Unix()%60)) * 50 + 50

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "# HELP app_requests_total Total HTTP requests\n")
	fmt.Fprintf(w, "# TYPE app_requests_total counter\n")
	fmt.Fprintf(w, "app_requests_total{method=\"GET\",endpoint=\"/\",status=\"200\"} %d\n", requestCount)

	fmt.Fprintf(w, "# HELP app_errors_total Total errors\n")
	fmt.Fprintf(w, "# TYPE app_errors_total counter\n")
	fmt.Fprintf(w, "app_errors_total{type=\"request_error\"} %d\n", errorCount)

	fmt.Fprintf(w, "# HELP app_processed_items Number of items processed\n")
	fmt.Fprintf(w, "# TYPE app_processed_items gauge\n")
	fmt.Fprintf(w, "app_processed_items{type=\"type_a\"} %d\n", processedItemsA)
	fmt.Fprintf(w, "app_processed_items{type=\"type_b\"} %d\n", processedItemsB)

	fmt.Fprintf(w, "# HELP app_cpu_load Application CPU load simulation\n")
	fmt.Fprintf(w, "# TYPE app_cpu_load gauge\n")
	fmt.Fprintf(w, "app_cpu_load %.2f\n", cpuLoad)

	fmt.Fprintf(w, "# HELP app_up Application uptime\n")
	fmt.Fprintf(w, "# TYPE app_up gauge\n")
	fmt.Fprintf(w, "app_up 1\n")
}

func main() {
	recordMetrics()

	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "healthy")
	})
	http.HandleFunc("/metrics", metricsHandler)

	fmt.Println("Starting sample app on :8888")
	http.ListenAndServe(":8888", nil)
}
