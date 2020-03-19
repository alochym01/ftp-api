package controllers

import (
	"fmt"
	"net/http"

	"github.com/alochym01/ftp-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AccountCreateInput validate request
type AccountCreateInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// AccountCcheckInput validate request
type AccountCheckInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password"`
	Mode     string `form:"mode"`
	Do       string `form:"do"`
}

// IsActive method chaining using gorm scope
func IsActive(db *gorm.DB) *gorm.DB {
	return db.Where("active = ?", 1)
}

// AccountCreate create an account
func AccountCreate(c *gin.Context) {
	var temp AccountCreateInput

	err := c.ShouldBind(&temp)
	if err != nil {
		fmt.Println("cannot bind request object")
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  false,
			"message": "missing username/password",
		})
		return
	}

	// Get database connection from GIN framework
	db := c.MustGet("db").(*gorm.DB)
	hashpassword, _ := models.HashPassword(temp.Password)
	acc := models.Account{Username: temp.Username, Password: hashpassword}

	result := db.Create(&acc)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": "Account is duplicated",
		})
		return
	}

	server := models.Server{}
	db.Where("Active = ?", 1).First(&server)

	nextserver := models.Server{}
	errs := db.Where("id = ?", server.ID+1).First(&nextserver)
	if errs != nil {
		db.Where("id = ?", 1).First(&nextserver)
	}

	// set active to next record
	server.Active = 0
	db.Save(&server)
	nextserver.Active = 1
	db.Save(&nextserver)

	// response result
	c.JSON(http.StatusOK, gin.H{
		"message": "Account is created",
		"result":  true,
		"ftp":     server.Domain,
		"live":    server.Domain,
		"port":    server.Port,
	})
}

// AccountDelete delete an account
func AccountDelete(c *gin.Context) {
	var temp AccountCreateInput

	err := c.ShouldBind(&temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing username/password",
			"result":  false,
		})
		return
	}

	// delete record in database
	db := c.MustGet("db").(*gorm.DB)
	acc := models.Account{}

	// db.Scopes(IsActive).Where("username = ? AND password = ?", temp.Username, hashpassword).Delete(models.Account{})
	err = db.Scopes(IsActive).Where("username = ?", temp.Username).First(&acc).Error
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Username is wrong",
			"result":  false,
		})
		return
	}

	// https://www.digitalocean.com/community/tutorials/understanding-boolean-logic-in-go
	if !models.CheckPasswordHash(temp.Password, acc.Password) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Password is wrong",
			"result":  false,
		})
		return
	}
	db.Delete(&acc)

	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"message": "Account is deleted",
	})
}

// AccountCheck check an account w user/pass and active
func AccountCheck(c *gin.Context) {
	var temp AccountCheckInput

	err := c.ShouldBind(&temp)
	if err != nil {
		c.String(http.StatusForbidden, "Username/Password is missing")
		return
	}

	// check record in database
	db := c.MustGet("db").(*gorm.DB)
	acc := models.Account{}

	err = db.Scopes(IsActive).Where("username = ?", temp.Username).First(&acc).Error
	if gorm.IsRecordNotFoundError(err) {
		c.String(http.StatusForbidden, "Username is wrong")
		return
	}

	if temp.Mode == "PAM_SM_AUTH" {
		if !models.CheckPasswordHash(temp.Password, acc.Password) {
			// https://www.digitalocean.com/community/tutorials/understanding-boolean-logic-in-go
			c.String(http.StatusForbidden, "Password is wrong")
			return
		}
	}

	c.String(http.StatusOK, "OK")
}
