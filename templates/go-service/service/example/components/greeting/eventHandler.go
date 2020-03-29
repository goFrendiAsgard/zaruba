package greeting

import (
	"registry.com/user/servicename/context"
	"registry.com/user/servicename/transport"
)

// CreateRegisterPersonHandler create event handler for "registerPerson"
func CreateRegisterPersonHandler(ctx *context.Context) transport.EventHandler {
	return func(msg transport.Message) (err error) {
		name, err := msg.GetString("name")
		if err != nil {
			return err
		}
		ctx.InitLocalCache("names", []string{})
		names, err := ctx.LocalCache.GetStringArray("names")
		if err != nil {
			return err
		}
		if name != "" {
			ctx.LocalCache["names"] = append(names, name)
		}
		return err
	}
}
