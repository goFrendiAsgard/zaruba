package example

import "github.com/gin-gonic/gin"

func getName(c *gin.Context) string {
	name := c.Param("name")
	if name == "" {
		name = c.Query("name")
	}
	if name == "" {
		name = c.PostForm("name")
	}
	return name
}
