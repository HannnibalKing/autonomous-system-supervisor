package bench

import (
	"context"
	"testing"
	"time"

	"autonomous-system-supervisor/pkg/decision"
)

func BenchmarkDecisionLoop(b *testing.B) {
	ctx := context.Background()
	engine := decision.NewEngine(1337)
	hc := decision.HealthContractStub("bench-component")
	for i := 0; i < b.N; i++ {
		_ = engine.Evaluate(ctx, hc, nil, decision.ClassificationStub())
	}
}

func BenchmarkDecisionLatency(b *testing.B) {
	ctx := context.Background()
	engine := decision.NewEngine(2025)
	hc := decision.HealthContractStub("bench-latency")
	start := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = engine.Evaluate(ctx, hc, nil, decision.ClassificationStub())
	}
	b.ReportMetric(float64(time.Since(start).Microseconds())/float64(b.N), "decision_us")
}
