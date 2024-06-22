package command

import "time"

type ParallelConfig struct {
	DryRun   bool
	Halt     *HaltConfig
	Jobs     uint
	Delay    time.Duration
	Progress bool
	Retries  uint
	Timeout  time.Duration
	Command  []string
}

type ParallelOption func(*ParallelConfig) error

func DryRun() ParallelOption {
	return func(cfg *ParallelConfig) error {
		cfg.DryRun = true

		return nil
	}
}

func WithHaltConfig(halt *HaltConfig) ParallelOption {
	return func(cfg *ParallelConfig) error {
		cfg.Halt = halt

		return nil
	}
}

func WithHalt(s string) ParallelOption {
	return func(cfg *ParallelConfig) error {
		halt, err := ParseHaltConfig(s)
		if err != nil {
			return err
		}

		cfg.Halt = &halt

		return nil
	}
}

func WithJobs(jobs uint) ParallelOption {
	return func(cfg *ParallelConfig) error {
		cfg.Jobs = jobs

		return nil
	}
}

func WithDelay(delay time.Duration) ParallelOption {
	return func(cfg *ParallelConfig) error {
		cfg.Delay = delay

		return nil
	}
}

func UpdateProgress() ParallelOption {
	return func(cfg *ParallelConfig) error {
		cfg.Progress = true

		return nil
	}
}

func WithRetries(retries uint) ParallelOption {
	return func(cfg *ParallelConfig) error {
		cfg.Retries = retries

		return nil
	}
}

func WithTimeout(timeout time.Duration) ParallelOption {
	return func(cfg *ParallelConfig) error {
		cfg.Timeout = timeout

		return nil
	}
}

func WithCommand(command []string) ParallelOption {
	return func(cfg *ParallelConfig) error {
		cfg.Command = command

		return nil
	}
}
