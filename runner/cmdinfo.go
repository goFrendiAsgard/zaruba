package runner

import (
	"io"
	"os/exec"
)

// CmdInfo represent information of Cmd
type CmdInfo struct {
	Cmd       *exec.Cmd
	IsProcess bool
	StdInPipe io.WriteCloser
	TaskName  string
}
