package controller

import (
	"net/http"

	"github.com/JRKs1532/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /typeproducts
func CreateTypeproduct(c *gin.Context) {
	var typeproduct entity.Typeproduct
	if err := c.ShouldBindJSON(&typeproduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&typeproduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": typeproduct})
}

// GET /typeproduct/:id
func GetTypeproduct(c *gin.Context) {
	var typeproduct entity.Typeproduct
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&typeproduct); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "typeproduct not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": typeproduct})
}

// GET /typeproducts
func ListTypeproducts(c *gin.Context) {
	var typeproducts []entity.Typeproduct
	if err := entity.DB().Raw("SELECT * FROM typeproducts").Scan(&typeproducts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": typeproducts})
}

// DELETE /typeproducts/:id
func DeleteTypeproduct(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM typeproducts WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "typeproduct not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /typeproducts
func UpdateTypeproduct(c *gin.Context) {
	var typeproduct entity.Typeproduct
	if err := c.ShouldBindJSON(&typeproduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", typeproduct.ID).First(&typeproduct); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "typeproduct not found"})
		return
	}

	if err := entity.DB().Save(&typeproduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": typeproduct})
}
