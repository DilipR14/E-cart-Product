package main

import (
	"log"
	"os"

	"github.com/DilipR14/E-cart-Products/controllers"
	"github.com/DilipR14/E-cart-Products/database"
	"github.com/DilipR14/E-cart-Products/middleware"
	"github.com/DilipR14/E-cart-Products/routes"
	"github.com/gin-gonic/gin"
)

// rest of your code...


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "2000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Product"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem()) // Fix: Change from app.removeItem() to app.RemoveItem()
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}