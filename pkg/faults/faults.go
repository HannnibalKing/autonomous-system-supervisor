package faults

// Classification describes whether a fault is transient or fatal.
type Classification struct {
	Transient bool
	Fatal     bool
	Label     string
}

// InjectorRequest captures a fault to inject for testing.
type InjectorRequest struct {
	TargetComponent string
	Type            string // network-partition, crash, time-skew, corruption
	Severity        string
}

// InjectionResult captures both the request and whether it was applied.
type InjectionResult struct {
	Request InjectorRequest
	Applied bool
}

// Classify uses a simple heuristic based on type.
func Classify(req InjectorRequest) Classification {
	switch req.Type {
	case "crash":
		return Classification{Fatal: true, Label: "crash"}
	case "network-partition":
		return Classification{Transient: true, Label: "network"}
	case "time-skew":
		return Classification{Transient: true, Label: "clock"}
	case "corruption":
		return Classification{Fatal: true, Label: "corruption"}
	default:
		return Classification{Label: "unknown"}
	}
}
