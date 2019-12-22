package command

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/state-alchemists/zaruba/format"
)

// Run a single command
func Run(dir, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	cmd.Dir, _ = filepath.Abs(dir)
	log.Printf("[INFO] Run `%s` on `%s` %s", command, cmd.Dir, format.SprintArgs(args))
	err := cmd.Run()
	return err
}
