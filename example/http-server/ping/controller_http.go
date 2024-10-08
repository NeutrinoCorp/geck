package ping

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/neutrinocorp/geck/observability/logging"
	"github.com/neutrinocorp/geck/systemerror"
	"github.com/neutrinocorp/geck/transport"
)

type ControllerHTTP struct {
	Logger logging.Logger
}

var _ transport.VersionedControllerHTTP = (*ControllerHTTP)(nil)

func NewControllerHTTP(logger logging.Logger) ControllerHTTP {
	return ControllerHTTP{
		Logger: logger,
	}
}

func (p ControllerHTTP) SetRoutes(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		p.Logger.Info().WriteWithCtx(c.Request().Context(), "got ping")
		return c.JSON(http.StatusOK, "pong")
	}, transport.IsResourceOwnerOrHasAnyAuthoritiesEcho("user_id", "ROLE_ADMIN"))
}

func (p ControllerHTTP) SetVersionedRoutes(g *echo.Group) {
	g.GET("/ping", func(c echo.Context) error {
		p.Logger.Info().WriteWithCtx(c.Request().Context(), "got versioned-ping")
		return systemerror.NewResourceNotFound[string]("foo")
	})
}
