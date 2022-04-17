// routes/routes.go
package routes

import (
	"rest-api-gostore/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/store", controllers.FindProducts)
	r.POST("/store", controllers.CreateProduct)
	r.GET("/store/:id", controllers.FindProduct)
	r.PATCH("/store/:id", controllers.UpdateProduct)
	r.DELETE("store/:id", controllers.DeleteProduct)

	r.GET("/category", controllers.FindCategories)
	r.POST("/category", controllers.CreateCategory)
	r.GET("/category/:id", controllers.FindCategory)
	r.PATCH("/category/:id", controllers.UpdateCategory)
	r.DELETE("category/:id", controllers.DeleteCategory)
	return r
}
