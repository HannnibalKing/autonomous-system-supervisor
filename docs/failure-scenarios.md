# Failure Scenarios

## Network Partition
- Agent loses connectivity to control plane; decision engine classifies as transient.
- Recovery: isolate workloads locally, buffer telemetry, retry with exponential backoff.

## Process Crash
- Agent or component crashes.
- Recovery: restart up to MaxRestart, otherwise migrate.

## Time Skew
- Node clock drifts beyond envelope.
- Recovery: mark telemetry unreliable, gate migrations, request NTP sync.

## Corrupted Messages
- Binary protocol checksum fails.
- Recovery: drop message, mark fault as fatal for the sender, request migration.

## Resource Exhaustion
- CPU/memory exceed safety envelope.
- Recovery: isolate and shed load; deterministic ordering ensures high-priority components recover first.
