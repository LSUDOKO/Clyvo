package routes

import (
	"github.com/LSUDOKO/Clyvo/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("users/signup", controller.SignUp())
	incommingRoutes.POST("users/login", controller.Login())
	incommingRoutes.POST("admin/addproduct", controller.ProductViewAdmin())
	incommingRoutes.Get("users/productview", controller.SearchProduct())
	incommingRoutes.Get("users/search", controller.SearchProductByQuery())

}
