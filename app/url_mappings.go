package app

import (
	"github.com/nytro04/asset_management/controllers/assets"
	"github.com/nytro04/asset_management/controllers/ping"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)

	router.GET("/assets/:id", assets.GetAsset)
	router.POST("/assets", assets.CreateAsset)
	router.PATCH("/assets/:id", assets.UpdateAsset)
	router.DELETE("/assets/:id", assets.DeleteAsset)

}