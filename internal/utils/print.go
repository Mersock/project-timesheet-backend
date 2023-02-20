package utils

import (
	"encoding/json"
	"log"
)

func Print(data interface{}) {
	manifestJson, _ := json.MarshalIndent(data, "", "  ")

	log.Println(string(manifestJson))
}
