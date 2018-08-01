package main

import (
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/require"
)

func TestPrint(t *testing.T) {
	fixtures := []struct {
		name   string
		args   []string
		stdout string
		stderr string
	}{{
		name:   "default",
		args:   []string{"test", "print"},
		stdout: "Message: my-message\n",
	}, {
		name:   "one_message",
		args:   []string{"test", "print", "--message", "hello"},
		stdout: "Message: hello\n",
	}}

	for _, fixture := range fixtures {
		t.Run(fixture.name, func(t *testing.T) {
			require := require.New(t)
			var (
				stdout, stderr string
				err            error
			)
			stdout = capturer.CaptureStdout(func() {
				stderr = capturer.CaptureStderr(func() {
					err = app.Run(fixture.args)
				})
			})

			require.NoError(err)
			require.Equal(fixture.stderr, stderr)
			require.Equal(fixture.stdout, stdout)
		})
	}
}
