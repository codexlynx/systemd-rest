package main

import (
	"github.com/codexlynx/systemd-rest/internal/core"
	"github.com/codexlynx/systemd-rest/internal/http/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	mode, address := core.GetConfiguration()
	if mode == core.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := routes.StartRouter()
	err := router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}
