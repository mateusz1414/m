package auth

import (
	"fmt"
	"matura/structs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func LoginUser(c *gin.Context) {
	user := structs.User{}
	c.BindJSON(&user)
	db, _ := c.Get("db")
	database := db.(*gorm.DB)
	err := user.Find(database)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Not authenticated",
		})
		return
	}
	token, err := user.GenerateToken()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"error": "Server problem",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "true",
		"email":   user.Email,
		"token":   token,
	})
}

func RegisterUser(c *gin.Context) {
	user := structs.User{}
	c.BindJSON(&user)
	db, _ := c.Get("db")
	database := db.(*gorm.DB)
	if user.Exist(database) {
		c.JSON(400, gin.H{
			"error": "User exist",
		})
		return
	}
	if err := user.Create(database); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "true",
	})
}
