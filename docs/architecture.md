# Autonomous System Supervisor Architecture

## Overview
A Go-based runtime supervisor combining control-plane logic with per-node agents. It monitors components via health contracts, enforces safety envelopes, classifies faults (transient vs fatal), and performs deterministic recovery and live migration. Observability and replay tooling ensure reproducibility.

## Components
- Control Plane: decision engine (rules + ML hook interface), policy manager, state store (pluggable), recovery planner.
- Agents: run on nodes, enforce health contracts, collect metrics, request migrations, isolate faults.
- Telemetry Plane: structured logs, metrics, traces; benchmark harness.
- Replay/Determinism: fixed seeds, binary dumps, deterministic timers.

## Data Flows
1. Agents emit health signals per contract.
2. Control plane applies decision engine and safety envelopes.
3. Fault classifier labels events; recovery planner orders actions deterministically.
4. Migrator executes state moves; telemetry captures metrics.
5. Replay tooling serializes state for deterministic re-run.

## Safety and Isolation
- Safety envelopes per component bound CPU/memory/latency.
- Fatal vs transient faults steer migrate vs restart paths.
- Recovery ordering uses stable sort to ensure deterministic execution.

## Why Go
- Concurrency (goroutines/channels) for low-latency supervision.
- Strong typing for contracts and protocols.
- Fast binaries and small footprint for agents.
- Testing and benchmarking built-in (`go test -bench`).
