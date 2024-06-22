package parallel

import (
	"context"
	"time"

	"github.com/bohdanch-w/parallel/command"
)

func NewFromConfig(cfg command.ParallelConfig) *Executor {
	return &Executor{}
}

type Executor struct {
	DryRun   bool
	Halt     *command.HaltConfig
	Jobs     uint
	Delay    uint
	Progress bool
	Retries  uint
	Timeout  time.Duration
	Command  []string
}

func Execute(ctx context.Context, _ command.ParallelConfig, args []string) error {
	return nil
}
