package main

import (
	"log"
	"net/http"

	"github.com/kotechi/api-comment/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	// Membuat koneksi ke database PostgreSQL
	dsn := "host=localhost user=postgres password=@pasific12op dbname=kubikitdb port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrasi model ke database
	db.AutoMigrate(&model.Comment{})

	// Menginisialisasi router Gin
	r := gin.Default()

	// Endpoint untuk mendapatkan semua komentar
	r.GET("/comments", getAllComments)

	// Endpoint untuk membuat komentar baru
	r.POST("/comments", createComment)

	// Endpoint untuk memperbarui komentar
	r.PUT("/comments/:id", updateComment)

	// Endpoint untuk menghapus komentar
	r.DELETE("/comments/:id", deleteComment)

	// Menjalankan server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// Handler untuk mendapatkan semua komentar
func getAllComments(c *gin.Context) {
	var comments []model.Comment
	if err := db.Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// Handler untuk membuat komentar baru
func createComment(c *gin.Context) {
	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, comment)
}

// Handler untuk memperbarui komentar
func updateComment(c *gin.Context) {
	var comment model.Comment
	id := c.Param("id")
	if err := db.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&comment)
	c.JSON(http.StatusOK, comment)
}

// Handler untuk menghapus komentar
func deleteComment(c *gin.Context) {
	var comment model.Comment
	id := c.Param("id")
	if err := db.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}
	db.Delete(&comment)
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
