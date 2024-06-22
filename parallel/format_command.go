package parallel

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

type commandFunc func(id uint, arg string) []string

func FormatCommandFunc(command []string, maxID uint) commandFunc {
	args := make([]commandArgFunc, 0, len(command))

	for _, arg := range command {
		args = append(args, formatCommandArgFunc(arg, maxID))
	}

	return func(id uint, arg string) []string {
		returnArgs := make([]string, 0, len(args))

		for _, argFunc := range args {
			returnArgs = append(returnArgs, argFunc(id, arg))
		}

		return returnArgs
	}
}

type commandArgFunc func(id uint, arg string) string

func formatCommandArgFunc(arg string, maxID uint) commandArgFunc {
	var (
		zeroPad = zeroPadIntFormat(maxID)
		parts   []argFormatter
		prev    int
	)

	for i := 0; i < len(arg); i++ {
		if arg[i] != '{' {
			continue
		}

		if i > prev {
			parts = append(parts, simpleString(arg[prev:i]))
		}

		switch {
		case strings.HasPrefix(arg[i:], "{}"):
			parts = append(parts, funcArgFormat(argNoChange))
			i++
		case strings.HasPrefix(arg[i:], "{.}"):
			parts = append(parts, funcArgFormat(argNoExtension))
			i += 2
		case strings.HasPrefix(arg[i:], "{/}"):
			parts = append(parts, funcArgFormat(argBase))
			i += 2
		case strings.HasPrefix(arg[i:], "{//}"):
			parts = append(parts, funcArgFormat(argDir))
			i += 3
		case strings.HasPrefix(arg[i:], "{/.}"):
			parts = append(parts, funcArgFormat(argBaseNoExt))
			i += 3
		case strings.HasPrefix(arg[i:], "{#}"):
			parts = append(parts, funcArgFormat(argID))
			i += 2
		case strings.HasPrefix(arg[i:], "{0#}"):
			parts = append(parts, argIDZeroPad(zeroPad))
			i += 3
		default:
			continue
		}

		prev = i + 1
	}

	if prev < len(arg)-1 {
		parts = append(parts, simpleString(arg[prev:]))
	}

	return func(id uint, arg string) string {
		var b strings.Builder

		for _, part := range parts {
			b.WriteString(part.format(id, arg))
		}

		return b.String()
	}
}

type argFormatter interface {
	format(id uint, arg string) string
}

type simpleString string

func (s simpleString) format(_ uint, _ string) string {
	return string(s)
}

type funcArgFormat func(uint, string) string

func (f funcArgFormat) format(id uint, arg string) string {
	return f(id, arg)
}

func argNoChange(_ uint, arg string) string {
	return arg
}

func argNoExtension(_ uint, arg string) string {
	return strings.TrimSuffix(arg, filepath.Ext(arg))
}

func argBase(_ uint, arg string) string {
	return filepath.Base(arg)
}

func argDir(_ uint, arg string) string {
	return filepath.ToSlash(filepath.Dir(arg))
}

func argBaseNoExt(_ uint, arg string) string {
	return argNoExtension(0, argBase(0, arg))
}

func argID(id uint, _ string) string {
	return strconv.Itoa(int(id))
}

func argIDZeroPad(format string) funcArgFormat {
	return func(id uint, _ string) string {
		return fmt.Sprintf(format, id)
	}
}

func zeroPadIntFormat(max uint) string {
	n := 1
	for max >= 10 {
		max /= 10
		n++
	}

	return fmt.Sprintf("%%0%dd", n)
}
