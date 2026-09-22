package main

import (
	"flag"
	"fmt"
	"os"

	"autonomous-system-supervisor/pkg/faults"
)

func main() {
	target := flag.String("target", "component-1", "component to fault")
	typ := flag.String("type", "network-partition", "fault type")
	severity := flag.String("severity", "low", "fault severity")
	report := flag.Bool("report", false, "sweep every known fault class and print a coverage report")
	flag.Parse()

	if *report {
		runCoverageReport(*target)
		return
	}

	req := faults.InjectorRequest{TargetComponent: *target, Type: *typ, Severity: *severity}
	class := faults.Classify(req)
	fmt.Printf("injected fault: %+v classification=%+v\n", req, class)
}

// runCoverageReport exercises every known (type, severity) fault class and
// reports what fraction produced a recognized classification. This is a
// regression-style safety metric: a drop in coverage means a fault class
// silently stopped classifying, not just that a single test failed.
func runCoverageReport(target string) {
	tracker := faults.NewCoverageTracker()
	for _, class := range faults.KnownFaultClasses() {
		req := faults.InjectorRequest{TargetComponent: target, Type: class.Type, Severity: class.Severity}
		tracker.Record(req, faults.Classify(req))
	}

	result := tracker.Report()
	fmt.Printf("fault-injection coverage: %d/%d classes (%.1f%%)\n", result.Covered, result.Total, result.Percent)
	for _, uncovered := range result.Uncovered {
		fmt.Printf("  UNCOVERED: type=%s severity=%s\n", uncovered.Type, uncovered.Severity)
	}

	if result.Percent < 100.0 {
		os.Exit(1)
	}
}
