package cli

import (
	"runtime"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/require"
)

func TestCompletionCommand(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("no completion command on windows")
	}

	require := require.New(t)
	app := New("test", "0.1.0", "abcde", "test app")
	var (
		stdout, stderr string
		err            error
	)

	stdout = capturer.CaptureStdout(func() {
		stderr = capturer.CaptureStderr(func() {
			err = app.Run([]string{"test", "completion"})
		})
	})

	require.NoError(err)
	require.Empty(stderr)
	require.Equal(
		`# Save this file to /etc/bash_completion.d/test
#
# or add the following line to your .bashrc file: 
#   echo "source <(test completion)" >> ~/.bashrc

_completion-test() {
    # All arguments except the first one
    args=("${COMP_WORDS[@]:1:$COMP_CWORD}")

    # Only split on newlines
    local IFS=$'\n'

    # Call completion (note that the first element of COMP_WORDS is
    # the executable itself)
    COMPREPLY=($(GO_FLAGS_COMPLETION=1 ${COMP_WORDS[0]} "${args[@]}"))
    return 0
}

complete -F _completion-test test
`, stdout)
}

func TestCompletionHelpCommand(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("no completion command on windows")
	}

	require := require.New(t)
	app := New("test", "0.1.0", "abcde", "test app")
	var (
		stdout, stderr string
		err            error
	)

	stdout = capturer.CaptureStdout(func() {
		stderr = capturer.CaptureStderr(func() {
			err = app.Run([]string{"test", "completion", "--help"})
		})
	})

	require.NoError(err)
	require.Empty(stderr)
	require.Equal(
			`Usage:
  test [OPTIONS] completion

Print a bash completion script for test.

You can place it on /etc/bash_completion.d/test, or add it to your .bashrc:
echo "source <(test completion)" >> ~/.bashrc


Help Options:
  -h, --help      Show this help message

`, stdout)
}
