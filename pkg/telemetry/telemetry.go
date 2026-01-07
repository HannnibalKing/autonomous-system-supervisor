package telemetry

import (
	"sync"
	"time"
)

// Metrics provides counters and timers for observability.
type Metrics interface {
	IncCounter(name string)
	ObserveLatency(name string, d time.Duration)
	Snapshot() map[string]float64
}

// InMemoryMetrics is a thread-safe metrics collector for tests and demos.
type InMemoryMetrics struct {
	mu       sync.Mutex
	counters map[string]float64
}

// NewInMemoryMetrics constructs a metrics sink.
func NewInMemoryMetrics() *InMemoryMetrics {
	return &InMemoryMetrics{counters: make(map[string]float64)}
}

func (m *InMemoryMetrics) IncCounter(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.counters[name]++
}

func (m *InMemoryMetrics) ObserveLatency(name string, d time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.counters[name] = d.Seconds()
}

func (m *InMemoryMetrics) Snapshot() map[string]float64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	copy := make(map[string]float64, len(m.counters))
	for k, v := range m.counters {
		copy[k] = v
	}
	return copy
}
