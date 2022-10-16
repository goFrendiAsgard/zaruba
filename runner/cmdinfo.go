package runner

import (
	"io"
	"os/exec"
	"time"

	"github.com/state-alchemists/zaruba/dsl"
)

// CmdInfo represent information of Cmd
type CmdInfo struct {
	Cmd                *exec.Cmd
	IsProcess          bool
	StdInPipe          io.WriteCloser
	Task               *dsl.Task
	CmdMaker           func() (*exec.Cmd, error) // function to recreate Cmd
	Attempt            int
	MaxRetry           int
	RetryDelayDuration time.Duration
}
