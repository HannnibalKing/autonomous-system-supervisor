package decision

import (
	"context"
	"math/rand"
	"time"

	"autonomous-system-supervisor/pkg/contracts"
	"autonomous-system-supervisor/pkg/faults"
)

// Result encodes the action to take for a component.
type Result struct {
	ComponentID string
	Action      string // restart, migrate, isolate, noop
	Reason      string
	Fatal       bool
}

// Engine executes rules and optionally ML hooks to determine actions.
type Engine struct {
	rng *rand.Rand
}

// NewEngine builds a deterministic engine from a fixed seed.
func NewEngine(seed int64) *Engine {
	return &Engine{rng: rand.New(rand.NewSource(seed))}
}

// HealthContractStub provides a deterministic contract for demos.
func HealthContractStub(id string) contracts.HealthContract {
	return contracts.HealthContract{ComponentID: id, HeartbeatFreq: 500 * time.Millisecond, Deadline: time.Second, MaxRestart: 3}
}

// ClassificationStub returns a default non-fatal classification.
func ClassificationStub() faults.Classification {
	return faults.Classification{Transient: true, Label: "demo"}
}

// Evaluate decides what to do given health status and faults.
func (e *Engine) Evaluate(ctx context.Context, hc contracts.HealthContract, breach *contracts.EnvelopeBreach, classification faults.Classification) Result {
	// Simple deterministic policy placeholder: if fatal or breaches, migrate; if transient, restart.
	switch {
	case classification.Fatal || (breach != nil && breach.Fatal):
		return Result{ComponentID: hc.ComponentID, Action: "migrate", Reason: "fatal-fault", Fatal: true}
	case classification.Transient:
		return Result{ComponentID: hc.ComponentID, Action: "restart", Reason: "transient-fault", Fatal: false}
	case breach != nil:
		return Result{ComponentID: hc.ComponentID, Action: "isolate", Reason: "envelope-breach", Fatal: breach.Fatal}
	default:
		_ = ctx
		// Controlled jitter for load shedding decisions.
		if e.rng.Float64() < 0.1 {
			return Result{ComponentID: hc.ComponentID, Action: "noop", Reason: "stable", Fatal: false}
		}
		return Result{ComponentID: hc.ComponentID, Action: "restart", Reason: "default-restart", Fatal: false}
	}
}
