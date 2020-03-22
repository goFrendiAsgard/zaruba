package pubsubhandlers

import (
	"registry.com/user/serviceName/communication"
	"registry.com/user/serviceName/servicedesc"
)

// CreateHelloHandler create hello handler for pubsub
func CreateHelloHandler(context *servicedesc.Context) communication.PubSubHandler {
	return func(input communication.Message) (err error) {

		// get name
		name := input["name"].(string)
		context.Logger.Printf("[RMQ PUBSUB] Hello %s", name)

		// add name to localCache
		context.InitLocalCache("names", []string{})
		oldNames, err := context.LocalCache.GetStringArray("names")
		if err != nil {
			return err
		}
		context.LocalCache["names"] = append(oldNames, name)

		return err
	}
}
