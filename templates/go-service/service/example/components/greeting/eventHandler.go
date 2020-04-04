package greeting

import (
	"app/context"
	"app/transport"
)

// CreateRegisterPersonHandler create event handler for "registerPerson"
func CreateRegisterPersonHandler(ctx *context.Context) transport.EventHandler {
	return func(msg transport.Message) (err error) {
		name, err := msg.GetString("name")
		if err != nil {
			return err
		}
		ctx.InitLocalCache("names", []string{})
		if name != "" {
			names, err := ctx.LocalCache.GetStringArray("names")
			if err != nil {
				return err
			}
			ctx.LocalCache["names"] = append(names, name)
		}
		return err
	}
}
