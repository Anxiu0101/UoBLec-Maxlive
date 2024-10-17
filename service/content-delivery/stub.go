package main

import (
	"Maxlive/conf"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Service structure to hold service name and address
type Service struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Port string `json:"port"`
}

var services []Service

func loadServices() ([]Service, error) {
	var services []Service

	// Open and read services.json
	file, err := os.Open("services.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode JSON into services slice
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&services)
	if err != nil {
		return nil, err
	}

	return services, nil
}

func main() {
	r := gin.Default()

	// Endpoint to receive requests and echo back
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": ServiceName})
	})

	// Load services
	services, err := conf.LoadServices()
	if err != nil {
		fmt.Println("Error loading services:", err)
		return
	}

	// Endpoint to get service addresses
	r.GET("/services", func(c *gin.Context) {
		c.JSON(http.StatusOK, services)
	})

	// Run server on default port
	r.Run(":8001")
}
