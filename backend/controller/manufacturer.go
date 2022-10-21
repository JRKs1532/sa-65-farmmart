package controller

import (
	"net/http"

	"github.com/JRKs1532/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /manufacturers
func CreateManufacturer(c *gin.Context) {
	var manufacturer entity.Manufacturer
	if err := c.ShouldBindJSON(&manufacturer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&manufacturer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": manufacturer})
}

// GET /manufacturer/:id
func GetManufacturer(c *gin.Context) {
	var manufacturer entity.Manufacturer
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&manufacturer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "manufacturer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": manufacturer})
}

// GET /manufacturers
func ListManufacturers(c *gin.Context) {
	var manufacturers []entity.Manufacturer
	if err := entity.DB().Raw("SELECT * FROM manufacturers").Scan(&manufacturers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": manufacturers})
}

// DELETE /manufacturers/:id
func DeleteManufacturer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM manufacturers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "manufacturer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /manufacturers
func UpdateManufacturer(c *gin.Context) {
	var manufacturer entity.Manufacturer
	if err := c.ShouldBindJSON(&manufacturer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", manufacturer.ID).First(&manufacturer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "manufacturer not found"})
		return
	}

	if err := entity.DB().Save(&manufacturer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": manufacturer})
}
