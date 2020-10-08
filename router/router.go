package router

import (
	"ShowWeb/controller"
	"ShowWeb/middleware"
	"github.com/gin-gonic/gin"
)

func RouterCollection(r *gin.Engine) *gin.Engine {
	r.Static("/static_file","./static_file/")

	r.POST("/register",controller.RegisterHandler)
	r.POST("/login",controller.LoginHandler)
	r.GET("/getmenu",controller.ReturnMenu)
	r.GET("/getplug",controller.GetPlugByMid)
	r.GET("/getmenubyid",controller.GetMenuByMid)
	r.GET("/gethotpbyid",controller.GetPlugsByMidHot)
	r.GET("/getdownloadadd",controller.DownloadAdd)
	r.GET("/getplugbykeyword",controller.GetPlugByKeyWord)
	r.POST("/createmenu",middleware.JWTAuth(),controller.CreateOneMenu)
	r.POST("/createplug",middleware.JWTAuth(),controller.CreateOnePlug)
	r.POST("/updatemenu",controller.UpdateOneMenu)
	r.POST("/updateplug",controller.UpdateOnePlug)
	r.POST("/deletemenu",controller.DeleteOneMenu)
	r.POST("/deleteplug",controller.DeleteOnePlug)

	return r
}
