package main

import (
	"github.com/LSUDOKO/Clyvo/controller"
	"github.com/LSUDOKO/Clyvo/databases"
	"github.com/LSUDOKO/Clyvo/middlewares"
	"github.com/LSUDOKO/Clyvo/routes"
	"github.com/gin-gonic/gin"

	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controller.NewApplication(databases.ProductData(databases.Client, "Products"), databases.UserData(databases.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middlewares.Authentication())
	router.GET("/addtocarts", app.AddToCart())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
