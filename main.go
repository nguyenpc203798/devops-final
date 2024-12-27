package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const PORT = "4000"

// Product struct để lưu thông tin sản phẩm
type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Slice để lưu trữ sản phẩm
var products = []Product{}

func main() {
	r := gin.Default()

	// Test endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Nguyen dep trai nhat the gioi",
		})
	})

	// CRUD endpoints
	r.POST("/products", createProduct)
	r.GET("/products", getProducts)
	r.GET("/products/:id", getProductByID)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)

	r.Run(":" + PORT)
}

// Tạo sản phẩm mới
func createProduct(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

// Lấy danh sách sản phẩm
func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

// Lấy sản phẩm theo ID
func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

// Cập nhật sản phẩm
func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product

	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, product := range products {
		if product.ID == id {
			updatedProduct.ID = id
			products[i] = updatedProduct
			c.JSON(http.StatusOK, updatedProduct)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

// Xóa sản phẩm
func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}
