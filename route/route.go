package route

import (
	"net/http"

	"github.com/butaosuinu/arma_mission_stockyard/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Routing() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	apiv1 := e.Group("/api/v1")

	apiv1.GET("/missions", controller.GetAllMissions)
	apiv1.GET("/missions/:id", controller.GetMission)
	apiv1.POST("/missions", controller.PostMission)
	apiv1.PUT("/missions/:id", controller.PutMission)
	apiv1.GET("/missions/search", func(c echo.Context) error {
		return c.String(http.StatusOK, "search res")
	})

	apiv1.GET("/makers/:id/missions", func(c echo.Context) error {
		return c.String(http.StatusOK, "makers")
	})

	return e
}
