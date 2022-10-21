package main

import (
	"github.com/JRKs1532/sa-65-example/controller"
	"github.com/JRKs1532/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	router := gin.Default()
	router.Use(CORSMiddleware())

	// Employee Routes
	router.GET("/employees", controller.ListEmployees)
	//router.GET("/employee/:id", controller.ListEmployees)
	router.GET("/employee/:id", controller.GetEmployee)
	router.POST("/employees", controller.CreateEmployee)
	router.PATCH("/employees", controller.UpdateEmployee)
	router.DELETE("/employees/:id", controller.DeleteEmployee)

	// Typeproduct Routes
	router.GET("/typeproducts", controller.ListTypeproducts)
	router.GET("/typeproduct/:id", controller.GetTypeproduct)
	router.POST("/typeproducts", controller.CreateTypeproduct)
	router.PATCH("/typeproducts", controller.UpdateTypeproduct)
	router.DELETE("/typeproducts/:id", controller.DeleteTypeproduct)

	// Manufacturer Routes
	router.GET("/manufacturers", controller.ListManufacturers)
	router.GET("/manufacturer/:id", controller.GetManufacturer)
	router.POST("/manufacturers", controller.CreateManufacturer)
	router.PATCH("/manufacturers", controller.UpdateManufacturer)
	router.DELETE("/manufacturers/:id", controller.DeleteManufacturer)

	// Product Routes
	router.GET("/products", controller.ListProducts)
	router.GET("/product/:id", controller.GetProduct)
	router.POST("products", controller.CreateProduct)
	router.PATCH("/products", controller.UpdateProduct)
	router.DELETE("/products/:id", controller.DeleteProduct)
	// Run the server

	router.Run()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
