package lib

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/abhishekghoshhh/gms/pkg/model"
)

const (
	RESOURCES                            = "resources"
	CAPABILITIES_FILE                    = "capabilities.xml"
	CAPABILITIES_CONFIG_REGEX_EXPRESSION = `\$\{(\w+)\}`
)

type CapabilityBuilder interface {
	Capabilities() string
}

type DefaultCapabilityBuilder struct {
	capabilities string
}

func (capabilityBuilder *DefaultCapabilityBuilder) Capabilities() string {
	return capabilityBuilder.capabilities
}
func CapabilitiesBuilder(capabilitiesConfig model.CapabilitiesConfig) *DefaultCapabilityBuilder {
	template := load(capabilitiesConfig)
	return &DefaultCapabilityBuilder{
		capabilities: template,
	}
}

func load(capabilitiesConfig model.CapabilitiesConfig) string {
	workingDir, _ := os.Getwd()
	capabilitiesDir := filepath.Join(workingDir, RESOURCES, CAPABILITIES_FILE)
	body, err := os.ReadFile(capabilitiesDir)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
		panic(err)
	}
	templateString := string(body)
	// Compile the regex pattern to match variables in the format ${variableName}
	re := regexp.MustCompile(CAPABILITIES_CONFIG_REGEX_EXPRESSION)

	// Use ReplaceAllStringFunc to replace each match with the corresponding value from the map
	result := re.ReplaceAllStringFunc(templateString, func(match string) string {
		// Extract the variable name from the match
		key := re.ReplaceAllString(match, "$1")
		// Return the value from the map or the original match if the variable is not found
		if value, ok := capabilitiesConfig.Get(key); ok {
			return value
		}
		return match
	})
	return result
}
