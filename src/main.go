package main

import (
	"github.com/codexlynx/systemd-rest/src/core"
	"github.com/codexlynx/systemd-rest/src/http/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	mode, address := core.GetConfiguration()
	if mode == core.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := routes.StartRouter()
	err := router.Run(address)
	if err != nil {
		panic(err)
	}
}
