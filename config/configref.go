package config

type ConfigRef struct {
	Map          map[string]string
	fileLocation string
	name         string
}

func (c *ConfigRef) GetFileLocation() string {
	return c.fileLocation
}

func (c *ConfigRef) GetName() string {
	return c.name
}
