package runner

// TaskStatus represent task status
type TaskStatus struct {
	Finished bool
	Error    error
}

// NewTaskStatus create new task status
func NewTaskStatus() (ts *TaskStatus) {
	return &TaskStatus{
		Finished: false,
		Error:    nil,
	}
}

// Finish task status
func (ts *TaskStatus) Finish(err error) {
	ts.Finished = true
	ts.Error = err
}
