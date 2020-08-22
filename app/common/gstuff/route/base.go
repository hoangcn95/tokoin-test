package route

import (
	"app/common/config"
	"app/common/gstuff/handler"
	gmiddleware "app/common/gstuff/middleware"

	"github.com/labstack/echo/v4"
)

var cfg = config.GetConfig()

func baseRoute(e *echo.Echo) *echo.Group {
	return e.Group("", gmiddleware.LogBody)
}

// APIRoute ..
func APIRoute(e *echo.Echo) *echo.Group {
	base := baseRoute(e)
	apiRoute := base.Group("/api")
	apiRoute.Any("/health", handler.Health)
	return apiRoute
}

// PublicAPIRoute ..
func PublicAPIRoute(e *echo.Echo) *echo.Group {
	base := baseRoute(e)
	return base.Group("/public-api")
}
