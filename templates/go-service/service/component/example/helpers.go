package example

// Context is interface of ginContext
type Context interface {
	Param(key string) string
	Query(key string) string
	PostForm(key string) string
}

func getName(c Context) string {
	name := c.Param("name")
	if name == "" {
		name = c.Query("name")
	}
	if name == "" {
		name = c.PostForm("name")
	}
	return name
}
