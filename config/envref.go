package config

type EnvRef struct {
	Map          map[string]Env
	Project      *Project
	fileLocation string
	name         string
}
