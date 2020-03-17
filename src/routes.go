package src

import (
	"fmt"
	"io"
	"os"

	"github.com/alochym01/ftp-api/src/controllers"
	"github.com/alochym01/ftp-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// InitRouter init all
func InitRouter() *gin.Engine {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Create table using GORM
	db.AutoMigrate(&models.Account{}, &models.Server{})

	// Create GIN engine framework
	r := gin.Default()

	// Provide gin controller db connection
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// fmt.Println(gin.Mode())
	// gin.SetMode(gin.ReleaseMode)
	fmt.Println(gin.Mode())

	// Home controller
	r.GET("/", controllers.Home)

	// Account controller
	acc := r.Group("/account")
	{
		acc.POST("/create", controllers.AccountCreate)
		acc.POST("/delete", controllers.AccountDelete)
		acc.POST("/check", controllers.AccountCheck)
	}
	// // cannot use Method=DELETE with GIN
	// r.DELETE("/account/delete", controllers.AccountDelete)

	// Server controller
	server := r.Group("/server")
	{
		server.POST("/create", controllers.ServerCreate)
	}

	return r
}
