package model

type (
	// Service structure to hold service name and address
	Service struct {
		Name string `json:"name"`
		Addr string `json:"addr"`
		Port string `json:"port"`
	}
)
