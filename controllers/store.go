package controllers

import (
	"net/http"
	"rest-api-gostore/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateStoreInput struct {
	Name        string `json:"name"`
	CodeProduct string `json:"codeProduct"`
	Price       string `json:"price"`
	Stock       string `json:"stock"`
}

type CreateCategoryInput struct {
	CodeCategory string `json:"codeCategory"`
	Name         string `json:"name"`
}

type UpdateStoreInput struct {
	Name        string `json:"name"`
	CodeProduct string `json:"codeProduct"`
	Price       string `json:"price"`
	Stock       string `json:"stock"`
}

type UpdateCategoryInput struct {
	CodeCategory string `json:"codeCategory"`
	Name         string `json:"name"`
}

// GET /Product
// Get all Product
func FindProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var stores []models.Product
	db.Find(&stores)

	c.JSON(http.StatusOK, gin.H{"data": stores})
}

// POST /Product
// Create new Product
func CreateProduct(c *gin.Context) {
	// Validate input
	var input CreateStoreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Product
	store := models.Product{Name: input.Name, Price: input.Price, Stock: input.Stock, CodeProduct: input.CodeProduct}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&store)

	c.JSON(http.StatusOK, gin.H{"data": store})
}

// GET /Product/:id
// Find a Product
func FindProduct(c *gin.Context) { // Get model if exist
	var store models.Product

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&store).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": store})
}

// PATCH /Product/:id
// Update a Product
func UpdateProduct(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var store models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&store).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateStoreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Product
	updatedInput.Name = input.Name
	updatedInput.CodeProduct = input.CodeProduct
	updatedInput.Price = input.Price
	updatedInput.Stock = input.Stock

	db.Model(&store).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": store})
}

// DELETE /Product/:id
// Delete a Product
func DeleteProduct(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var store models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&store).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&store)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

/* ## Controller Category ##*/

// GET /Category
// Get all Category
func FindCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var categories []models.Category
	db.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// POST /Category
// Create new Category
func CreateCategory(c *gin.Context) {
	// Validate input
	var input CreateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Category
	store := models.Category{Name: input.Name, CodeCategory: input.CodeCategory}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&store)

	c.JSON(http.StatusOK, gin.H{"data": store})
}

// GET /Category/:id
// Find a Category
func FindCategory(c *gin.Context) { // Get model if exist
	var categories models.Category

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// PATCH /Categorys/:id
// Update a Category
func UpdateCategory(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var categories models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Category
	updatedInput.CodeCategory = input.CodeCategory
	updatedInput.Name = input.Name

	db.Model(&categories).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// DELETE /Category/:id
// Delete a Category
func DeleteCategory(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var categories models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&categories)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
