package controllers

import (
	"github.com/codexlynx/systemd-rest/internal/core"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUnitJournal(ctx *gin.Context) {
	unitName := ctx.Params.ByName("name")
	switch journal, err := core.ReadUnitJournal(unitName); err.(type) {
	case *core.InvalidUnitName:
		ctx.Status(http.StatusNotFound)
	case nil:
		ctx.Data(http.StatusOK, "text/plain", journal)
	default:
		log.Print(err)
	}
	return
}
