package controllers

import (
	"github.com/codexlynx/systemd-rest/src/core"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUnits(ctx *gin.Context) {
	units, err := core.GetUnits()
	if err != nil {
		log.Print(err)
	}
	ctx.JSON(http.StatusOK, units)
	return
}

func GetUnit(ctx *gin.Context) {
	unitName := ctx.Params.ByName("name")
	switch unit, err := core.GetUnit(unitName); err.(type) {
	case *core.InvalidUnitName:
		ctx.Status(http.StatusNotFound)
	case nil:
		ctx.JSON(http.StatusOK, unit)
	default:
		log.Print(err)
	}
	return
}

func StartUnit(ctx *gin.Context) {
	unitName := ctx.Params.ByName("name")
	_, wait := ctx.GetQuery("wait")
	switch err := core.StartUnit(unitName, wait); err.(type) {
	case *core.InvalidUnitName:
		ctx.Status(http.StatusNotFound)
	case nil:
		ctx.Status(http.StatusCreated)
	default:
		log.Print(err)
	}
	return
}

func StopUnit(ctx *gin.Context) {
	unitName := ctx.Params.ByName("name")
	_, wait := ctx.GetQuery("wait")
	switch err := core.StopUnit(unitName, wait); err.(type) {
	case *core.InvalidUnitName:
		ctx.Status(http.StatusNotFound)
	case nil:
		ctx.Status(http.StatusCreated)
	default:
		log.Print(err)
	}
	return
}
