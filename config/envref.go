package config

type EnvRef struct {
	BaseEnvMap   map[string]Env
	Project      *Project
	fileLocation string
	name         string
}
