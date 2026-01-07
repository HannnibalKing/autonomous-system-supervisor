package main

import (
	"flag"
	"fmt"

	"autonomous-system-supervisor/pkg/faults"
)

func main() {
	target := flag.String("target", "component-1", "component to fault")
	typ := flag.String("type", "network-partition", "fault type")
	severity := flag.String("severity", "low", "fault severity")
	flag.Parse()

	req := faults.InjectorRequest{TargetComponent: *target, Type: *typ, Severity: *severity}
	class := faults.Classify(req)
	fmt.Printf("injected fault: %+v classification=%+v\n", req, class)
}
