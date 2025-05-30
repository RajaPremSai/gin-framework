package middleware

import (
	"fmt"
	"gin-framework/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckMiddleware(c *gin.Context){
	headers := c.GetHeader("Authorization")

	fmt.Println("Headers:", headers)
	if headers == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	token := strings.Split(headers, "Bearer ")

	if len(token) < 2 {
		c.JSON(401, gin.H{"error": "Unauthorized, token not provided"})
		c.Abort()
		return
	}

	err := utils.TokenCheck(token[1])
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized, invalid token"})
		c.Abort()
		return
	}

	c.Next()
}