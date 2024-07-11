package index

import (
	"fmt"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.SecureJSON(http.StatusOK, gin.H{"message": "Hello world!"})
}

func HelloHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.SecureJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var data map[string]string
	if err := c.BindJSON(&data); err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	username, ok := data["username"]
	if !ok {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Missing 'username' in request body"})
		return
	}

	response := fmt.Sprintf("Hello %s", username)
	c.SecureJSON(http.StatusOK, gin.H{"msg": response})
}
