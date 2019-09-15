package main

import (
	"fmt"
	"log"
	"os"

	"github.com/butaosuinu/arma_mission_stockyard/infra"
	"github.com/butaosuinu/arma_mission_stockyard/route"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "dev")
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	infra.Init()

	e := route.Routing()

	e.Start(":8080")
}
