package runner

import (
	"os/exec"
	"sync"

	"github.com/state-alchemists/zaruba/modules/config"
)

type processState struct {
	executed bool
	process  *exec.Cmd
}

func createProcessState(process *exec.Cmd) *processState {
	return &processState{
		executed: false,
		process:  process,
	}
}

// Runner stateful
type Runner struct {
	stoppedLock        *sync.RWMutex
	processesStateLock *sync.RWMutex
	stopped            bool
	processState       map[string]*processState
	p                  *config.ProjectConfig
	componentsToRun    map[string]*config.Component
}

// CreateRunner create runner
func CreateRunner(p *config.ProjectConfig, selectors []string) (r *Runner, err error) {
	components, err := p.GetComponentsBySelectors(selectors)
	if err != nil {
		return r, err
	}
	r = &Runner{
		p:               p,
		stopped:         false,
		processState:    map[string]*processState{},
		componentsToRun: components,
	}
	return r, err
}

func (r *Runner) stop() {
	r.stoppedLock.Lock()
	r.stopped = true
	r.stoppedLock.Unlock()
}

func (r *Runner) getStopped() bool {
	r.stoppedLock.RLock()
	stopped := r.stopped
	r.stoppedLock.RUnlock()
	return stopped
}

func (r *Runner) stopByChan(stopChan chan bool) {
	<-stopChan
	r.stop()
}

func (r *Runner) registerProcess(name string, process *exec.Cmd) {

}

// Run run components
func (r *Runner) Run(projectDir string, stopChan, executedChan chan bool, errChan chan error) {
	go r.stopByChan(stopChan)
}
