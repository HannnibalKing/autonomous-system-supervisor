# Autonomous System Supervisor (Self-Healing Orchestrator)

Go runtime supervisor for autonomous systems: monitors components, enforces safety envelopes, classifies faults, restarts/isolates/migrates workloads, and supports deterministic replay.

## Layout
- cmd/supervisor: control-plane demo entrypoint.
- cmd/agent: node agent demo entrypoint.
- pkg/*: contracts, decision engine, faults, migration hooks, recovery ordering, telemetry, replay.
- bench: Go benchmarks for decision loop.
- tools/fault-injector: CLI for network partition/crash/time-skew/corruption injection classification.
- docs: architecture, failure scenarios, design tradeoffs, benchmarks.

## Quickstart
```
go run ./cmd/supervisor

go run ./cmd/agent

go test ./...

go test ./bench -bench . -benchmem
```

## Determinism & Replay
- Fixed seeds for decision engine stubs.
- Stable recovery ordering.
- Replay package supports binary seed dumps for deterministic reruns.

## Benchmarks
See docs/benchmarks.md and run the `bench` package benchmarks for latency/throughput/memory/goroutine counts.
