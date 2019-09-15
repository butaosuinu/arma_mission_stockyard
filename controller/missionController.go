package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type msg struct {
	Message string `json:"message"`
}

func GetMission(c echo.Context) (err error) {
	missionID := c.Param("id")
	id, err := strconv.Atoi(missionID)
	if err != nil {
		log.Println(err)
		return c.JSON(
			http.StatusBadRequest,
			msg{Message: "ID must be number"},
		)
	}
	fmt.Println(id)
	return c.JSON(http.StatusOK, nil)
}

func GetAllMissions(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, nil)

}

func PostMission(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, nil)

}

func PutMission(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, nil)

}
