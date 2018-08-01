package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/require"
)

func TestSleep(t *testing.T) {
	fixtures := []struct {
		signal     os.Signal
		signalName string
	}{{
		signalName: "SIGTERM",
		signal:     syscall.SIGTERM,
	}, {
		signalName: "SIGINT",
		signal:     syscall.SIGINT,
	}}

	for _, fixture := range fixtures {
		t.Run(fixture.signalName, func(t *testing.T) {
			require := require.New(t)
			var (
				stdout, stderr string
				err            error
			)

			if fixture.signal != nil {
				go func() {
					time.Sleep(1 * time.Second)
					pid := os.Getpid()
					p, err := os.FindProcess(pid)
					if err != nil {
						panic(err)
					}

					p.Signal(fixture.signal)
				}()
			}

			stdout = capturer.CaptureStdout(func() {
				stderr = capturer.CaptureStderr(func() {
					err = app.Run([]string{"test", "sleep"})
				})
			})

			require.NoError(err)
			require.True(strings.Contains(stdout, "Sleeping...\n"))
			if fixture.signal != nil {
				require.True(strings.Contains(stderr,
					fmt.Sprintf("signal %s received", fixture.signalName)))
			}
		})
	}
}
