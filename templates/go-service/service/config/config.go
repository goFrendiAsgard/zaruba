package config

// Config is a general service context
type Config struct {
	HTTPPort                  int
	ServiceName               string
	GlobalRmqConnectionString string
	LocalRmqConnectionString  string
}

// CreateConfig initiate new config
func CreateConfig() (config *Config) {
	return &Config{
		HTTPPort:                  GetIntFromEnv("SERVICENAME_HTTP_PORT", 3000),
		ServiceName:               "servicename",
		GlobalRmqConnectionString: GetStrFromEnv("GLOBAL_RMQ_CONNECTION_STRING", "amqp://localhost:15672/"),
		LocalRmqConnectionString:  GetStrFromEnv("LOCAL_RMQ_CONNECTION_STRING", "amqp://localhost:15672/"),
	}
}
