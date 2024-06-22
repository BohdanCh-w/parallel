package parallel

// --dry-run
// --halt
// --jobs
// --arg-file
// --delay
// --delimiter
// --progress
// --no-run-if-empty (default: true) // TODO: implement
// --retries n
// --timeout

// {}
// {.}
// {/}
// {//}
// {/.}
// {#}
// {0#} -- custom

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/bohdanch-w/parallel/command"
)

type cmdConfig struct {
	command.ParallelConfig
	argFile   string
	delimiter byte
	timeout   time.Duration
}

func (cfg *cmdConfig) parseArgs(args []string) error {
	var (
		fset            = flag.NewFlagSet("parallel", flag.ExitOnError)
		halt, delimiter string
	)

	fset.BoolVar(&cfg.DryRun, "dry-run", false, "Print the command that would be executed, but do not execute it.")
	fset.StringVar(&halt, "halt", "", "Stop the execution if the specified condition is met.")
	fset.UintVar(&cfg.Jobs, "jobs", 1, "The number of jobs to run concurrently.")
	fset.StringVar(&cfg.argFile, "arg-file", "", "Read arguments from the specified file.")
	fset.DurationVar(&cfg.Delay, "delay", 0, "Delay between each command.")
	fset.StringVar(&delimiter, "delimiter", "\n", "The delimiter to use when splitting the input.")
	fset.BoolVar(&cfg.Progress, "progress", false, "Show progress.")
	fset.UintVar(&cfg.Retries, "retries", 0, "The number of times to retry a command.")
	fset.DurationVar(&cfg.Timeout, "timeout", 0, "The maximum time a command is allowed to run.")
	fset.DurationVar(&cfg.timeout, "total-timeout", 0, "The maximum time all commands are allowed to run.")

	if err := fset.Parse(args); err != nil {
		return fmt.Errorf("parse flags: %w", err)
	}

	if halt != "" {
		haltCond, err := command.ParseHaltConfig(halt)
		if err != nil {
			return fmt.Errorf("parse halt: %w", err)
		}

		cfg.Halt = &haltCond
	}

	if len(delimiter) != 1 {
		return fmt.Errorf("delimiter must be a single character")
	}

	cfg.delimiter = delimiter[0]
	cfg.Command = fset.Args()

	return cfg.validate()
}

func (cfg *cmdConfig) validate() error {
	if cfg.Jobs == 0 {
		return fmt.Errorf("jobs must be greater than 0")
	}

	if len(cfg.Command) == 0 || strings.Join(cfg.Command, "") == "" {
		return fmt.Errorf("command must not be empty")
	}

	return nil
}
