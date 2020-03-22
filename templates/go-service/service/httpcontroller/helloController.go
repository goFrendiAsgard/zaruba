package httpcontroller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Hello send classical hello world or hello + name
func Hello(c *gin.Context) {
	// get name
	name := c.Param("name")
	if name == "" {
		name = c.Query("name")
	}
	if name == "" {
		name = c.PostForm("name")
	}
	if name == "" {
		name = "world"
	}
	// send response
	c.String(200, fmt.Sprintf("Hello %s !!!", name))
}
