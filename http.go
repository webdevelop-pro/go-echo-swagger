package echoswagger

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/go-openapi/loads"
	echo "github.com/labstack/echo/v4"

	swagMW "github.com/go-openapi/runtime/middleware"
	"github.com/webdevelop-pro/go-common/configurator"
)

const pkgName = "echoswagger"

func New(c *configurator.Configurator, e *echo.Echo) error {
	cfg := c.New(pkgName, &Config{}, pkgName).(*Config)

	specDoc, err := loads.Spec("./" + cfg.FILE_PATH)
	if err != nil {
		return err
	}

	doc, err := json.MarshalIndent(specDoc.Spec(), "", "  ")
	if err != nil {
		return err
	}

	e.Use(echo.WrapMiddleware(
		func(h http.Handler) http.Handler {
			return swagMW.Spec("/", doc, h)
		},
	))

	e.Use(echo.WrapMiddleware(
		func(h http.Handler) http.Handler {
			return swagMW.SwaggerUI(swagMW.SwaggerUIOpts{
				BasePath: "/",
				SpecURL:  path.Join("/", cfg.FILE_PATH[0:len(cfg.FILE_PATH)]),
				Path:     cfg.URL_PATH,
			}, h)
		},
	))

	return nil
}
