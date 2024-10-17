package main

import (
	"encoding/json"
	"fmt"
	"io"
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

// Find service by name
func getServiceAddrByName(serviceName string) (string, error) {
	for _, service := range services {
		if service.Name == serviceName {
			path := fmt.Sprintf("%s:%s", service.Addr, service.Port)
			return path, nil
		}
	}
	return "", fmt.Errorf("service not found")
}

func main() {
	r := gin.Default()

	// Endpoint to receive requests and echo back
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Load services
	services, err := loadServices()
	if err != nil {
		fmt.Println("Error loading services:", err)
		return
	}

	// Endpoint to get service addresses
	r.GET("/services-list", func(c *gin.Context) {
		c.JSON(http.StatusOK, services)
	})

	// Endpoint to forward request to a service
	r.GET("/api/v1/:service", func(c *gin.Context) {
		serviceName := c.Param("service")

		// Get the address of the service
		serviceAddr, err := getServiceAddrByName(serviceName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
			return
		}

		// Forward the request to the target service
		pingpath := fmt.Sprintf("http://%s/ping", serviceAddr)
		println(pingpath)
		resp, err := http.Get(pingpath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to forward request"})
			return
		}
		defer resp.Body.Close()

		// Read the response from the target service
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from service"})
			return
		}

		// Return the response from the target service to the client
		c.Data(resp.StatusCode, "application/json", respBody)
	})

	// Run server on default port
	err = r.Run(":8000")
	if err != nil {
		fmt.Println("Error running services:", err)
		return
	}
}
