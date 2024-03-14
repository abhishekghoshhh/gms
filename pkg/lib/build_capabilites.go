package lib

import (
	"log"
	"os"
)

var template string

func init() {
	body, err := os.ReadFile("resources/capabilities.xml")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
		panic(err)
	}
	template = string(body)
}

func CapabilitiesTemplate() string {
	return template
}
