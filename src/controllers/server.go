package controllers

import (
	"net/http"

	"github.com/alochym01/ftp-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ServerInfo request validate
type ServerInfo struct {
	Domain string `form:"domain" binding:"required"`
	Port   uint   `form:"port" binding:"required"`
}

// ServerCreate save info
func ServerCreate(c *gin.Context) {
	var s ServerInfo

	err := c.ShouldBind(&s)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  false,
			"message": "missing domain/port",
			"error":   err.Error(),
		})
		return
	}

	// Get database from GIN framework
	db := c.MustGet("db").(*gorm.DB)

	server := models.Server{}
	errs := db.Where("domain = ? AND port = ?", s.Domain, s.Port).First(&server)

	if errs != nil {
		// for _, e := range errs.GetErrors() {
		// 	fmt.Println(e)
		// }
		if gorm.IsRecordNotFoundError(errs.Error) {
			server := models.Server{Domain: s.Domain, Port: s.Port}
			db.Create(&server)
			// check server is first record of server table
			if server.ID == 1 {
				server.Active = 1
				db.Save(&server)
			}
			c.JSON(http.StatusOK, gin.H{
				"result": "Domain is created",
			})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Domain/Port is duplicated",
	})
}
