package faults

// KnownFaultClass is one (type, severity) combination the injector supports.
type KnownFaultClass struct {
	Type     string
	Severity string
}

// KnownFaultClasses lists every fault class the fault-injector is expected to exercise.
// Extend this when a new fault type or severity is added to the injector.
func KnownFaultClasses() []KnownFaultClass {
	types := []string{"network-partition", "crash", "time-skew", "corruption"}
	severities := []string{"low", "medium", "high"}
	classes := make([]KnownFaultClass, 0, len(types)*len(severities))
	for _, t := range types {
		for _, s := range severities {
			classes = append(classes, KnownFaultClass{Type: t, Severity: s})
		}
	}
	return classes
}

// CoverageTracker records which known fault classes have been exercised with a
// classified, non-unknown result. This is a leading-indicator metric: it answers
// "how much of our known fault space has a verified recovery classification?"
// rather than a simple test pass/fail count.
type CoverageTracker struct {
	exercised map[KnownFaultClass]bool
}

// NewCoverageTracker creates an empty tracker.
func NewCoverageTracker() *CoverageTracker {
	return &CoverageTracker{exercised: make(map[KnownFaultClass]bool)}
}

// Record marks a fault class as exercised if it classified to a known label.
func (c *CoverageTracker) Record(req InjectorRequest, class Classification) {
	if class.Label == "unknown" {
		return
	}
	c.exercised[KnownFaultClass{Type: req.Type, Severity: req.Severity}] = true
}

// CoverageReport summarizes exercised vs known fault classes.
type CoverageReport struct {
	Total     int
	Covered   int
	Percent   float64
	Uncovered []KnownFaultClass
}

// Report computes coverage against KnownFaultClasses.
func (c *CoverageTracker) Report() CoverageReport {
	known := KnownFaultClasses()
	report := CoverageReport{Total: len(known)}
	for _, class := range known {
		if c.exercised[class] {
			report.Covered++
		} else {
			report.Uncovered = append(report.Uncovered, class)
		}
	}
	if report.Total > 0 {
		report.Percent = 100.0 * float64(report.Covered) / float64(report.Total)
	}
	return report
}
