package controller

import (
	"net/http"

	"github.com/JRKs1532/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /products
func CreateProduct(c *gin.Context) {

	var product entity.Product
	var typeproduct entity.Typeproduct
	var manufacturer entity.Manufacturer
	var employee entity.Employee

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา typeproduct ด้วย id
	if tx := entity.DB().Where("id = ?", product.TypeproductID).First(&typeproduct); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "typeproduct not found"})
		return
	}

	// 10: ค้นหา manufacturer ด้วย id
	if tx := entity.DB().Where("id = ?", product.ManufacturerID).First(&manufacturer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "manufacturer not found"})
		return
	}

	// 11: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", product.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	// 12: สร้าง Product
	wv := entity.Product{
		Product_name:  product.Product_name,  // ตั้งค่าฟิลด์ Product_name
		Product_price: product.Product_price, //ตั้งค่าฟิลด์ Product_price
		Typeproduct:   typeproduct,           // โยงความสัมพันธ์กับ Entity Typeproduct
		Manufacturer:  manufacturer,          // โยงความสัมพันธ์กับ Entity Typeproduct
		Employee:      employee,              // โยงความสัมพันธ์กับ Entity Employee

	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": wv})
}

// GET /product/:id
func GetProduct(c *gin.Context) {
	var product entity.Product
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GET /products
func ListProducts(c *gin.Context) {
	var products []entity.Product
	if err := entity.DB().Preload("Typeproduct").Preload("Manufacturer").Preload("Employee").Raw("SELECT * FROM products").Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// DELETE products/:id
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM products WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /products
func UpdateProduct(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", product.ID).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	if err := entity.DB().Save(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}
