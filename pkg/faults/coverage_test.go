package faults

import "testing"

func TestCoverageTracker_FullSweepIsFullyCovered(t *testing.T) {
	tracker := NewCoverageTracker()
	for _, class := range KnownFaultClasses() {
		req := InjectorRequest{TargetComponent: "component-1", Type: class.Type, Severity: class.Severity}
		tracker.Record(req, Classify(req))
	}

	report := tracker.Report()
	if report.Percent != 100.0 {
		t.Fatalf("expected 100%% coverage after exercising every known class, got %.2f%% uncovered=%v",
			report.Percent, report.Uncovered)
	}
}

func TestCoverageTracker_UnknownTypeDoesNotCount(t *testing.T) {
	tracker := NewCoverageTracker()
	req := InjectorRequest{TargetComponent: "component-1", Type: "bogus-fault", Severity: "low"}
	tracker.Record(req, Classify(req))

	report := tracker.Report()
	if report.Covered != 0 {
		t.Fatalf("expected an unclassified fault type to contribute zero coverage, got %d", report.Covered)
	}
}

func TestCoverageTracker_PartialSweepReportsUncovered(t *testing.T) {
	tracker := NewCoverageTracker()
	req := InjectorRequest{TargetComponent: "component-1", Type: "crash", Severity: "high"}
	tracker.Record(req, Classify(req))

	report := tracker.Report()
	if report.Covered != 1 {
		t.Fatalf("expected exactly one covered class, got %d", report.Covered)
	}
	if len(report.Uncovered) != report.Total-1 {
		t.Fatalf("expected %d uncovered classes, got %d", report.Total-1, len(report.Uncovered))
	}
}
