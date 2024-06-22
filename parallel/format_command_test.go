package parallel_test

import (
	"strings"
	"testing"

	"github.com/bohdanch-w/parallel/parallel"
)

func TestFormatCommandFunc(t *testing.T) {
	testcases := []struct {
		name     string
		args     []string
		maxID    uint
		id       uint
		arg      string
		expected string
	}{
		{
			name:     "simple",
			args:     []string{"echo", "foo"},
			arg:      "abc",
			expected: "echo foo",
		},
		{
			name:     "no change",
			args:     []string{"echo", "{}"},
			arg:      "foo",
			expected: "echo foo",
		},
		{
			name:     "no extension",
			args:     []string{"echo", "{.}"},
			arg:      "foo.txt",
			expected: "echo foo",
		},
		{
			name:     "base",
			args:     []string{"echo", "{/}"},
			arg:      "/path/to/foo.txt",
			expected: "echo foo.txt",
		},
		{
			name:     "dir",
			args:     []string{"echo", "{//}"},
			arg:      "/path/to/foo.txt",
			expected: "echo /path/to",
		},
		{
			name:     "base no extension",
			args:     []string{"echo", "{/.}"},
			arg:      "/path/to/foo.txt",
			expected: "echo foo",
		},
		{
			name:     "id",
			args:     []string{"echo", "{#}"},
			id:       42,
			arg:      "foo",
			expected: "echo 42",
		},
		{
			name:     "id zero padded",
			args:     []string{"echo", "{0#}"},
			maxID:    100,
			id:       42,
			arg:      "foo",
			expected: "echo 042",
		},
		{
			name:     "multiple",
			args:     []string{"echo", "{#}", "{/}"},
			id:       42,
			arg:      "/path/to/foo.txt",
			expected: "echo 42 foo.txt",
		},
		{
			name:     "multiple in one",
			args:     []string{"echo", "{#} {/}"},
			id:       42,
			arg:      "/path/to/foo.txt",
			expected: "echo 42 foo.txt",
		},
		{
			name:     "multiple in one + extra",
			args:     []string{"echo", "prefix_{#}_infix_{/}_suffix"},
			id:       42,
			arg:      "/path/to/foo.txt",
			expected: "echo prefix_42_infix_foo.txt_suffix",
		},
		{
			name:     "multiple in one no space",
			args:     []string{"echo", "{#}{/}"},
			id:       42,
			arg:      "/path/to/foo.txt",
			expected: "echo 42foo.txt",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := parallel.FormatCommandFunc(tc.args, tc.maxID)(tc.id, tc.arg)

			if s := strings.Join(actual, " "); s != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, s)
			}
		})
	}
}
