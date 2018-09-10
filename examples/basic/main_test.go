package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	cmd := exec.Command("go", "build", "-o", "test", ".")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func TestMain(t *testing.T) {
	require := require.New(t)

	cmd := exec.Command("./test", "--help")

	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()
	require.NoError(err)

	require.Empty(stderr.String())
	require.True(strings.HasPrefix(stdout.String(), "Usage:\n  basic"))
}
