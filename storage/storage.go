package storage

import (
    "sync"
    "time"
)

type DataPoint struct {
    Timestamp time.Time `json:"timestamp"`
    Value     float64   `json:"value"`
}

type MemoryStorage struct {
    mu      sync.RWMutex
    metrics map[string][]DataPoint
}

func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{
        metrics: make(map[string][]DataPoint),
    }
}

func (s *MemoryStorage) AddMetric(name string, ts time.Time, value float64) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.metrics[name] = append(s.metrics[name], DataPoint{Timestamp: ts, Value: value})
}

func (s *MemoryStorage) GetAll() map[string][]DataPoint {
    s.mu.RLock()
    defer s.mu.RUnlock()
    copy := make(map[string][]DataPoint)
    for k, v := range s.metrics {
        copy[k] = append([]DataPoint(nil), v...)
    }
    return copy
}

func (s *MemoryStorage) GetMetric(name string) []DataPoint {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return append([]DataPoint(nil), s.metrics[name]...)
}