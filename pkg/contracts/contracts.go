package contracts

import "time"

// HealthContract describes health signals and thresholds for a component.
type HealthContract struct {
	ComponentID   string
	HeartbeatFreq time.Duration
	Deadline      time.Duration
	MaxRestart    int
}

// SafetyEnvelope defines safety limits for a component to operate within.
type SafetyEnvelope struct {
	ComponentID string
	MaxCPU      float64
	MaxMemoryMB int
	MaxLatency  time.Duration
}

// EnvelopeBreach captures envelope violations to classify severity.
type EnvelopeBreach struct {
	ComponentID string
	Metric      string
	Observed    float64
	Threshold   float64
	Fatal       bool
}
