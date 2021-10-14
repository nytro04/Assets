package assets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nytro04/asset_management/domain/assets"
	"github.com/nytro04/asset_management/services"
	"github.com/nytro04/asset_management/utils/errors"
)

//handles every request thats coming in to the application
// checks if the request has everything needed to process it and
// sends it to the services

func CreateAsset(ctx *gin.Context) {
	var asset assets.Asset
	if err := ctx.ShouldBindJSON(&asset); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateAsset(asset)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
func GetAsset(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")

}
func UpdateAsset(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")

}
func DeleteAsset(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")

}
