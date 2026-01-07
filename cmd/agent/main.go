package main

import (
	"context"
	"log"
	"time"

	"autonomous-system-supervisor/pkg/agent"
	"autonomous-system-supervisor/pkg/contracts"
	"autonomous-system-supervisor/pkg/faults"
	"autonomous-system-supervisor/pkg/telemetry"
)

func main() {
	ctx := context.Background()
	metrics := telemetry.NewInMemoryMetrics()
	a := agent.New("agent-1", metrics, 99)
	contract := contracts.HealthContract{ComponentID: "component-1", HeartbeatFreq: 300 * time.Millisecond, Deadline: time.Second, MaxRestart: 3}
	breach := &contracts.EnvelopeBreach{ComponentID: contract.ComponentID, Metric: "latency", Observed: 120.0, Threshold: 100.0, Fatal: false}
	fault := faults.InjectionResult{Request: faults.InjectorRequest{TargetComponent: contract.ComponentID, Type: "network-partition", Severity: "medium"}, Applied: true}
	res := a.Tick(ctx, contract, breach, fault)
	log.Printf("decision: %+v", res)
	log.Printf("metrics: %+v", metrics.Snapshot())
}
