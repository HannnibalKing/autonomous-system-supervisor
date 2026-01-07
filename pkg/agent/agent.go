package agent

import (
	"context"
	"time"

	"autonomous-system-supervisor/pkg/contracts"
	"autonomous-system-supervisor/pkg/decision"
	"autonomous-system-supervisor/pkg/faults"
	"autonomous-system-supervisor/pkg/telemetry"
)

// Agent executes health checks and enforces contracts on a node.
type Agent struct {
	id      string
	metrics telemetry.Metrics
	engine  *decision.Engine
}

// New creates an agent with deterministic decision-making.
func New(id string, metrics telemetry.Metrics, seed int64) *Agent {
	return &Agent{id: id, metrics: metrics, engine: decision.NewEngine(seed)}
}

// Tick runs a single health evaluation cycle.
func (a *Agent) Tick(ctx context.Context, contract contracts.HealthContract, breach *contracts.EnvelopeBreach, fault faults.InjectionResult) decision.Result {
	class := faults.Classify(fault.Request)
	start := time.Now()
	res := a.engine.Evaluate(ctx, contract, breach, class)
	a.metrics.IncCounter("agent.tick")
	a.metrics.ObserveLatency("agent.decision.latency", time.Since(start))
	return res
}
