# Mini Prometheus Clone (Golang)

This is a minimal monitoring tool built with Golang, inspired by Prometheus. It scrapes metrics from configured targets, stores them in memory, and exposes a REST API to retrieve metric data.

## 📦 Features

- Scrapes Prometheus-style metrics from HTTP targets
- Stores time-series data in-memory
- REST API to access all or specific metrics
- Dummy metrics target included
- Dockerized for easy deployment

## 🛠️ How to Use

### 1. Run the Dummy Exporter (target)

```bash
go run target/main.go
```

It serves metrics on `http://localhost:8081/metrics`.

### 2. Run Mini Prometheus

```bash
go run main.go
```

It starts scraping the dummy target and exposes an API on `http://localhost:9090`.

### 3. API Endpoints

- `GET /metrics` – all scraped metrics
- `GET /metrics/{name}` – specific metric by name

### 4. Docker

```bash
docker-compose up --build
```

## 📈 Example Metric Output

```
http_requests_total 5
cpu_temperature_celsius 65.3
```

---
