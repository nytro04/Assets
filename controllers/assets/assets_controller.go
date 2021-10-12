package assets

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//handles every request thats coming in to the application
// checks if the request has everything needed to process it and
// sends it to the services

func CreateAsset(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "implement me")

}
func GetAsset(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "implement me")

}
func UpdateAsset(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "implement me")

}
func DeleteAsset(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "implement me")

}