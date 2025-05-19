package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	fmt.Fprintf(w, "# TYPE http_requests_total counter\n")
	fmt.Fprintf(w, "http_requests_total %d\n", rand.Intn(100))
	fmt.Fprintf(w, "# TYPE cpu_temperature_celsius gauge\n")
	fmt.Fprintf(w, "cpu_temperature_celsius %.2f\n", 60+rand.Float64()*10)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/metrics", metricsHandler)
	fmt.Println("Dummy exporter running at :8081/metrics")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
