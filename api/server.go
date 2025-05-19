package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"mini-prometheus/storage"
)

func StartServer(store *storage.MemoryStorage) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, `<h1>Mini Prometheus</h1><p>Use <a href="/metrics">/metrics</a> to view scraped metrics.</p>`)
			return
		}

		if strings.HasPrefix(r.URL.Path, "/metrics") {
			metricsHandler(store, w, r)
			return
		}

		http.NotFound(w, r)
	})

	http.ListenAndServe(":9090", nil)
}

func metricsHandler(store *storage.MemoryStorage, w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) == 2 && parts[1] == "metrics" {
		metrics := store.GetAll()
		json.NewEncoder(w).Encode(metrics)
		return
	}

	if len(parts) == 3 && parts[1] == "metrics" {
		name := parts[2]
		data := store.GetMetric(name)
		json.NewEncoder(w).Encode(data)
		return
	}

	http.NotFound(w, r)
}
