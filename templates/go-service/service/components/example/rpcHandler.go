package example

import (
	"errors"
	"fmt"
)

// GreetRPCController greet with RPC
func GreetRPCController(inputs ...interface{}) (greeting interface{}, err error) {
	if len(inputs) == 0 {
		return greeting, errors.New("Message accepted but input is invalid")
	}
	name, success := inputs[0].(string)
	if !success {
		errorMessage := fmt.Sprintf("Cannot convert %#v to string", inputs[0])
		return greeting, errors.New(errorMessage)
	}
	return Greet(name), err
}
