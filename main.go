package main

import (
	"github.com/butaosuinu/arma_mission_stockyard/route"
)

func main() {
	e := route.Routing()

	e.Start(":8080")
}
