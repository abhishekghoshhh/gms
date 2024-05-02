package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	capabilites = `<?xml version="1.0" encoding="UTF-8"?>
<vosi:capabilities
	xmlns:vosi="http://www.ivoa.net/xml/VOSICapabilities/v1.0"
	xmlns:vs="http://www.ivoa.net/xml/VODataService/v1.1"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<capability standardID="ivo://ivoa.net/std/VOSI#capabilities">
		<interface xsi:type="vs:ParamHTTP" role="std">
			<accessURL use="full">http://localhost:8080/gms/capabilities</accessURL>
		</interface>
	</capability>
</vosi:capabilities>`
)

func TestCapabilities(t *testing.T) {
	workingDir, _ := os.Getwd()
	workingDir = filepath.Dir(workingDir)
	workingDir = filepath.Dir(workingDir)

	t.Run("should return capabilites in xml wih sucess code ok", func(t *testing.T) {
		config := map[string]string{
			"scheme":    "http",
			"proxyName": "localhost",
			"proxyPort": "8080",
		}

		capabilitiesPath := filepath.Join(
			workingDir,
			"resources/test/capabilities.xml",
		)

		handler := CapabilitiesHandler(config, capabilitiesPath)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/gms/search", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.GetTemplate(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/xml", rec.Header().Get("Content-Type"))
		assert.Equal(t, capabilites, rec.Body.String())
	})

	t.Run("should panic if the file is not present", func(t *testing.T) {
		config := map[string]string{
			"scheme":    "http",
			"proxyName": "localhost",
			"proxyPort": "8080",
		}

		capabilitiesPath := filepath.Join(
			workingDir,
			"resources/test/dummy-capabilities.xml",
		)

		defer func() {
			r := recover()
			assert.NotNil(t, r)
		}()
		CapabilitiesHandler(config, capabilitiesPath)
	})
}
