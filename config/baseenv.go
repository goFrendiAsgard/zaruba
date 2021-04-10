package config

type BaseEnv struct {
	From    string
	Default string
}

type ProjectBaseEnv struct {
	BaseEnvMap   map[string]BaseEnv
	Project      *Project
	fileLocation string
	name         string
}
