package routers

import (
	"backend/middleware"
	v1 "backend/servers/v1"
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
	// request headers allowed to exec
	config.AllowHeaders = []string{"tus-presumable", "upload-length", "upload-metadata", "cache-control", "x-requested-with", "*"}

	r.Use(cors.New(config))
	apiv1 := r.Group("/api/v1")
	// router division: some of the routers set JWTAuthMiddleware
	commodity := apiv1.Group("")
	{
		// home
		commodity.GET("home/", v1.Home)
		// commodity list
		commodity.GET("commodity/list/", v1.CommodityDetail)
		// user register/login
		commodity.POST("shopper/login/", v1.ShopperLogin)
	}
	shopper := apiv1.Group("", middleware.JWTAuthMiddleware)
	{
		// commodity collect
		shopper.POST("shopper/collect/", v1.CommodityCollect)
		// Logout
		shopper.POST("shopper/logout/", v1.ShopperLogout)
		// individual home
		shopper.GET("shopper/home/", v1.ShopperHome)
		// add into the cart
		shopper.POST("shopper/cart/", v1.ShopperShopCart)
		// Cart list
		shopper.GET("shopper/cart/", v1.ShopperShopCart)
		// online payment
		shopper.POST("shopper/pays/", v1.ShopperPays)
		// delete commodity in the cart
		shopper.POST("shopper/delete/", v1.ShopperDelete)
	}
	return r
}
