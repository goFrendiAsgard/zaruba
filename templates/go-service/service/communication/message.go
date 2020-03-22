package communication

import (
	"errors"
	"fmt"
)

// Message for RPC and pub-sub
type Message map[string]interface{}

// GetInterface  get interface property of Message
func (m Message) GetInterface(key string) (val interface{}, err error) {
	val, exists := m[key]
	if !exists {
		errorMessage := fmt.Sprintf("Key `%s` doesn't exist on %#v", key, val)
		err = errors.New(errorMessage)
	}
	return val, err
}

// GetInt64 of Message
func (m Message) GetInt64(key string) (val int64, err error) {
	data, err := m.GetInterface(key)
	if err != nil {
		return val, err
	}
	val, success := data.(int64)
	if !success {
		errorMessage := fmt.Sprintf("Conversion to int64 failed: %#v", val)
		err = errors.New(errorMessage)
	}
	return val, err
}

// GetInt32 of Message
func (m Message) GetInt32(key string) (val int32, err error) {
	data, err := m.GetInterface(key)
	if err != nil {
		return val, err
	}
	val, success := data.(int32)
	if !success {
		errorMessage := fmt.Sprintf("Conversion to int32 failed: %#v", val)
		err = errors.New(errorMessage)
	}
	return val, err
}

// GetFloat64 of Message
func (m Message) GetFloat64(key string) (val float64, err error) {
	data, err := m.GetInterface(key)
	if err != nil {
		return val, err
	}
	val, success := data.(float64)
	if !success {
		errorMessage := fmt.Sprintf("Conversion to float64 failed: %#v", val)
		err = errors.New(errorMessage)
	}
	return val, err
}

// GetFloat32 of Message
func (m Message) GetFloat32(key string) (val float32, err error) {
	data, err := m.GetInterface(key)
	if err != nil {
		return val, err
	}
	val, success := data.(float32)
	if !success {
		errorMessage := fmt.Sprintf("Conversion to float32 failed: %#v", val)
		err = errors.New(errorMessage)
	}
	return val, err
}

// GetInt of Message
func (m Message) GetInt(key string) (val int, err error) {
	data, err := m.GetInterface(key)
	if err != nil {
		return val, err
	}
	val, success := data.(int)
	if !success {
		errorMessage := fmt.Sprintf("Conversion to int failed: %#v", val)
		err = errors.New(errorMessage)
	}
	return val, err
}

// GetBool of Message
func (m Message) GetBool(key string) (val bool, err error) {
	data, err := m.GetInterface(key)
	if err != nil {
		return val, err
	}
	val, success := data.(bool)
	if !success {
		errorMessage := fmt.Sprintf("Conversion to bool failed: %#v", val)
		err = errors.New(errorMessage)
	}
	return val, err
}

// GetString of Message
func (m Message) GetString(key string) (val string, err error) {
	data, err := m.GetInterface(key)
	if err != nil {
		return val, err
	}
	val, success := data.(string)
	if !success {
		errorMessage := fmt.Sprintf("Conversion to string failed: %#v", val)
		err = errors.New(errorMessage)
	}
	return val, err
}

// GetStringArray of Message
func (m Message) GetStringArray(key string) (val []string, err error) {
	data, err := m.GetInterface(key)
	if err != nil {
		return val, err
	}
	val, success := data.([]string)
	if !success {
		errorMessage := fmt.Sprintf("Conversion to string array failed: %#v", val)
		err = errors.New(errorMessage)
	}
	return val, err
}
