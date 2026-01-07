package main

import (
	"context"
	"log"
	"time"

	"autonomous-system-supervisor/pkg/control"
	"autonomous-system-supervisor/pkg/telemetry"
)

func main() {
	ctx := context.Background()
	metrics := telemetry.NewInMemoryMetrics()
	supervisor := control.NewSupervisor(metrics, 42)
	if err := supervisor.Run(ctx, 200*time.Millisecond); err != nil {
		log.Fatal(err)
	}
	log.Printf("supervisor metrics: %+v", metrics.Snapshot())
}
