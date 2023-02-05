package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {

	var text string = "We are debbuging from Docker :D! Is iNotifyWait working?"

	fmt.Println(text)

	fmt.Println("Segundo breakpoint :D")

	c.JSON(http.StatusOK, gin.H{
		"API_VERSION": "v1",
		"Message":     "Hi! Your service is running as expected :D! Funcionando con Air y delve",
	})
}
