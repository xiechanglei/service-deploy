package service

import (
	"encoding/json"
	"log"
	"os"
)

type Service struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Url   int    `json:"url"`
	Port  int    `json:"port"`
	Dir   string `json:"dir"`
	Cmd   string `json:"cmd"`
	Async bool   `json:"async"`
	Delay int    `json:"delay"`
}

func ReadService() []Service {
	filePath := "./service.json"
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("could not read file:  service.json")
		return nil
	}
	var services []Service
	err = json.Unmarshal(contentBytes, &services)
	if err != nil {
		log.Println("could not unmarshal service.json", err)
		return nil
	}
	return services
}
