package control

import (
	"context"
	"time"

	"autonomous-system-supervisor/pkg/decision"
	"autonomous-system-supervisor/pkg/recovery"
	"autonomous-system-supervisor/pkg/telemetry"
)

// Supervisor coordinates decisions and recovery ordering across components.
type Supervisor struct {
	metrics telemetry.Metrics
	engine  *decision.Engine
}

// NewSupervisor constructs a supervisor with deterministic seed.
func NewSupervisor(metrics telemetry.Metrics, seed int64) *Supervisor {
	return &Supervisor{metrics: metrics, engine: decision.NewEngine(seed)}
}

// Run drives a small control loop for demonstration.
func (s *Supervisor) Run(ctx context.Context, tick time.Duration) error {
	ticker := time.NewTicker(tick)
	defer ticker.Stop()
	plan := recovery.NewPlan([]recovery.Step{{ComponentID: "c1", Priority: 1, Action: "restart"}})
	_ = plan
	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			start := time.Now()
			// Placeholder: evaluate a fake component.
			_ = s.engine.Evaluate(ctx, decision.HealthContractStub("c1"), nil, decision.ClassificationStub())
			s.metrics.IncCounter("supervisor.tick")
			s.metrics.ObserveLatency("supervisor.loop.latency", time.Since(start))
		}
	}
	return nil
}
