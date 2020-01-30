package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tarathep/hellogoapi/api"
)

//Route is routing init
type Route struct {
	Handler api.HelloHandler
}

//SetupRouter is setup routing
func (h Route) SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", h.Handler.GetHello)
	r.POST("/hello", h.Handler.PostHello)
	return r
}
