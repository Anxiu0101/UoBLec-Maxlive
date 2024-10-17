package main

import (
	"Maxlive/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	services []model.Service
)

func main() {
	r := gin.Default()

	// Endpoint to receive requests and echo back
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": ServiceName})
	})

	////Load services
	//services, err := conf.LoadServices()
	//if err != nil {
	//	fmt.Println("Error loading services:", err)
	//	return
	//}

	//r.GET("/services-list", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, services)
	//})

	//// Endpoint to get service addresses
	//r.GET("/echo", service.EchoService)

	// Run server on default port
	r.Run(":8002")
}
