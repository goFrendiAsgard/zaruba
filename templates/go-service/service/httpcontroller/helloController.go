package httpcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
  
// ParamFetcher  ...  
type ParamFetcher func(key string) (val string)

// Hello send classical hello world or hello + name
func Hello(c *gin.Context) {
	for _, fetch := range([]ParamFetcher{c.Param, c.Query, c.PostForm}) {
		name := fetch("name")
		if name != "" {
			c.String(200, fmt.Sprintf("Hello %s !!!", name))
			return
		}
	}
	c.String(200, "Hello world !!!")
}
