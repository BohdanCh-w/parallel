package parallel

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
)

func RunParallel(ctx context.Context, args []string) error {
	var cfg cmdConfig

	if err := cfg.parseArgs(args); err != nil {
		return err
	}

	return run(ctx, cfg)
}

func run(ctx context.Context, cfg cmdConfig) error {
	if cfg.timeout > 0 {
		var cancel context.CancelFunc

		ctx, cancel = context.WithTimeout(ctx, cfg.timeout)
		defer cancel()
	}

	cmdArgs, err := parseCommandArgs(cfg.argFile, cfg.delimiter)
	if err != nil {
		return fmt.Errorf("parse command args: %w", err)
	}

	return nil
}

func pprintJSON(v any) {
	s := fmt.Sprintf("%+v", v)

	fmt.Println(s)

	return
}

func parseCommandArgs(argFile string, delimiter byte) ([]string, error) {
	var input io.Reader

	if argFile != "" {
		f, err := os.Open(argFile)
		if err != nil {
			return nil, fmt.Errorf("open arg file: %w", err)
		}

		defer f.Close()

		input = f
	} else {
		input = os.Stdin
	}

	var (
		reader = bufio.NewReader(input)
		args   []string
	)

	for {
		arg, err := reader.ReadString(delimiter)
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, fmt.Errorf("read input: %w", err)
		}

		args = append(args, arg)
	}

	return args, nil
}
