package routers

import (
	"backend/settings"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(settings.Mode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.StaticFS("/static", http.Dir("static"))
	// set cross-origin access
	config := cors.DefaultConfig()
	// allow all the origins
	config.AllowAllOrigins = true
	// methods allowed to exec
	config.AllowMethods = []string{"GET", "POST"}

	return r
}
