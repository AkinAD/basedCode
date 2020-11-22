package main

import (
	// auth "github.com/AkinAD/basedCode/auth"
	// item "github.com/AkinAD/basedCode/item"
	// user "github.com/AkinAD/basedCode/user"
	"github.com/gin-gonic/gin"
)

// the '/' endpoint
func homeHandler(c *gin.Context) {
	c.JSON(
		200,
		gin.H{"message": "hello"},
	)
}
