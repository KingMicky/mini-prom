# Build stage
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o mini-prometheus main.go

# Run stage
FROM gcr.io/distroless/base-debian10
COPY --from=builder /app/mini-prometheus /mini-prometheus
COPY config.yaml /config.yaml
ENTRYPOINT ["/mini-prometheus"]