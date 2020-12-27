package controllers

import (
	"github.com/codexlynx/systemd-rest/src/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUnitJournal(ctx *gin.Context)  {
	unitName := ctx.Params.ByName("name")
	switch content, err := core.ReadUnitJournal(unitName); err.(type) {
	default:
		panic(err)
	case *core.InvalidUnitName:
		ctx.Status(http.StatusNotFound)
	case nil:
		ctx.Data(http.StatusOK, "text/plain", content)
	}
	return
}
