package web

import (
	groute "app/common/gstuff/route"
	"app/web/handler"

	"github.com/labstack/echo/v4"
)

func route(e *echo.Echo) {
	API := groute.APIRoute(e)
	API.GET("/users/search", handler.UsersHandler.Search)
	API.GET("/tickets/search", handler.TicketsHandler.Search)
	API.GET("/organizations/search", handler.OrganizationsHandler.Search)
}
