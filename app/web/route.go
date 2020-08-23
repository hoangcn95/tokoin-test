package web

import (
	groute "app/common/gstuff/route"
	"app/web/handler"

	"github.com/labstack/echo/v4"
)

func route(e *echo.Echo) {
	API := groute.APIRoute(e)
	APIv1 := API.Group("/v1")
	APIv1.GET("/users/search", handler.UsersHandler.Search)
	APIv1.GET("/tickets/search", handler.TicketsHandler.Search)
	APIv1.GET("/organizations/search", handler.OrganizationsHandler.Search)
}
