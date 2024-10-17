package conf

import (
	"Maxlive/model"
	"encoding/json"
	"os"
)

func LoadServices() (services []model.Service, err error) {
	// Open and read services.json
	file, err := os.Open("content-creation\\services.json")
	if err != nil {
		println(err.Error())
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
