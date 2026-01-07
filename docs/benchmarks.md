# Benchmarks

Run `go test ./bench -bench . -benchmem`.

Metrics to capture:
- Latency: decision loop p50/p99.
- Throughput: health events per second processed.
- Memory: allocation per decision and steady-state heap.
- Goroutines: steady vs load.

Sample placeholders (replace with real runs):
- Decision loop: p50 12µs, p99 40µs.
- Throughput: 250k eval/s with 1 engine.
- Memory: <64 bytes/op allocation in benchmark harness.
- Goroutines: <40 during stress test.
