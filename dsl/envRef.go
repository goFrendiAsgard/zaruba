package dsl

type EnvRef struct {
	Map          map[string]*Env
	Project      *Project
	fileLocation string
	name         string
}

func (e *EnvRef) GetFileLocation() string {
	return e.fileLocation
}

func (e *EnvRef) GetName() string {
	return e.name
}
