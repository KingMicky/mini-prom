package main

import (
    "log"
    "mini-prometheus/config"
    "mini-prometheus/scraper"
    "mini-prometheus/api"
    "mini-prometheus/storage"
)

func main() {
    cfg, err := config.LoadConfig("config.yaml")
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    store := storage.NewMemoryStorage()
    for _, target := range cfg.Targets {
        go scraper.StartScraper(target.URL, target.Interval, store)
    }

    api.StartServer(store)
}