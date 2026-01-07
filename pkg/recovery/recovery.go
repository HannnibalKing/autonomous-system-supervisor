package recovery

import "sort"

// Step represents a deterministic recovery action.
type Step struct {
	ComponentID string
	Priority    int
	Action      string
}

// Plan orders recovery steps deterministically by priority then component id.
type Plan struct {
	Steps []Step
}

// NewPlan creates a deterministic ordering.
func NewPlan(steps []Step) Plan {
	sorted := make([]Step, len(steps))
	copy(sorted, steps)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Priority == sorted[j].Priority {
			return sorted[i].ComponentID < sorted[j].ComponentID
		}
		return sorted[i].Priority < sorted[j].Priority
	})
	return Plan{Steps: sorted}
}
