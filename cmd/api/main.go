package main

import (
	"github.com/CriticalNoob02/store-control-be/internal/route"
	"github.com/CriticalNoob02/store-control-be/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./infra/env/.env")
	util.Check(err)

	api := gin.Default()

	route.ClientRoutes(api)

	api.Run()
}
