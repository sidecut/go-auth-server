package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMethod(c *gin.Context) {
	fmt.Println("\n'GetMethod' called")
	IdValue := c.Params.ByName("IdValue")
	message := "GetMethod Called With Param: " + IdValue
	c.JSON(http.StatusOK, message)
	reqPayload := make([]byte, 1024)
	reqPayload, err := c.GetRawData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Request Payload Data: ", string(reqPayload))
}

func main() {
	router := gin.Default()
	subRouterAuthenticated := router.Group("/api/v1/PersonId", gin.BasicAuth(gin.Accounts{
		"admin": "adminpass",
	}))
	subRouterAuthenticated.GET("/:IdValue", GetMethod)
	listenPort := "1357"

	router.Run(":" + listenPort)
}
