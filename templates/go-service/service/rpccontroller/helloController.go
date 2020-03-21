package rpccontroller

import (
	"registry.com/user/serviceName/communication"
)

// Hello classical hello world or hello + name
func Hello(input communication.Message) (output communication.Message, err error) {
	name := input["name"].(string)
	output = communication.Message{
		"greeting": "Hello " + name,
	}
	return output, err
}
