package api

import (
	"net/http"
	"os"
	"regexp"

	"github.com/labstack/echo"
)

const (
	CAPABILITIES_CONFIG_REGEX_EXPRESSION = `\$\{(\w+)\}`
)

type CapabilitiesApi struct {
	capabilitesTemplate string
}

func CapabilitiesHandler(config map[string]string, filePath string) *CapabilitiesApi {
	template := loadCapabilities(config, filePath)
	return &CapabilitiesApi{
		capabilitesTemplate: template,
	}
}

func (handler *CapabilitiesApi) GetTemplate(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/xml")
	return c.String(http.StatusOK, handler.capabilitesTemplate)
}

func loadCapabilities(config map[string]string, filePath string) string {
	body, err := os.ReadFile(filePath)
	if err != nil {
		panic("unable to read file: " + err.Error())
	}
	templateString := string(body)
	re := regexp.MustCompile(CAPABILITIES_CONFIG_REGEX_EXPRESSION)

	result := re.ReplaceAllStringFunc(templateString, func(match string) string {
		key := re.ReplaceAllString(match, "$1")
		if value, ok := config[key]; ok {
			return value
		}
		return match
	})
	return result
}
