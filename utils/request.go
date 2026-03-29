package utils

import "github.com/gin-gonic/gin"

// BindJSON is a reusable helper for parsing JSON request bodies
func BindJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return err
	}
	return nil
}
