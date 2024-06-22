package command

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bohdanch-w/parallel/internal"
)

type HaltConfig struct {
	Completed   uint
	Failed      uint
	Succeeded   uint
	KillRunning bool
}

func ParseHaltConfig(s string) (HaltConfig, error) {
	var cfg HaltConfig

	runningAction, condition, _ := strings.Cut(s, ",")
	if runningAction == "" || condition == "" {
		return cfg, fmt.Errorf("%w: %q", internal.Error("invalid format"), s)
	}

	conditionType, conditionValueStr, _ := strings.Cut(condition, "=")
	if conditionType == "" || conditionValueStr == "" {
		return cfg, fmt.Errorf("%w: %q", internal.Error("invalid condition format"), s)
	}

	conditionValue, err := strconv.Atoi(conditionValueStr)
	if err != nil {
		return cfg, fmt.Errorf("%w: %q", internal.Error("invalid condition value"), s)
	}

	switch runningAction {
	case "now":
		cfg.KillRunning = true
	case "soon":
		cfg.KillRunning = false
	default:
		return cfg, fmt.Errorf("%w: %q", internal.Error("invalid running action"), s)
	}

	switch conditionType {
	case "completed":
		cfg.Completed = uint(conditionValue)
	case "failed":
		cfg.Failed = uint(conditionValue)
	case "succeeded":
		cfg.Succeeded = uint(conditionValue)
	default:
		return cfg, fmt.Errorf("%w: %q", internal.Error("invalid condition type"), s)
	}

	if conditionValue <= 0 {
		return cfg, internal.Error("halt condition value must be greater than 0")
	}

	return cfg, nil
}
