package math

import (
	"matura/structs"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AddFunction(c *gin.Context) {
	user, _ := c.Get("userID")
	userID := user.(int)
	db, _ := c.Get("db")
	database := db.(*gorm.DB)
	function := structs.Function{}
	c.BindJSON(&function)
	function.UserID = userID
	err := function.AddFunction(database)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Add function problem",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "true",
		"id":      function.FunctionID,
	})

}

func DeleteFunction(c *gin.Context) {
	functionID, _ := strconv.Atoi(c.Param("id"))
	user, _ := c.Get("userID")
	userID := user.(int)
	db, _ := c.Get("db")
	database := db.(*gorm.DB)
	function := structs.FindFunction(userID, functionID, database)
	if function.Value == "" {
		c.JSON(400, gin.H{
			"error": "Function not found",
		})
		return
	}
	err := function.RemoveFunction(database)
	if err != nil {

		c.JSON(400, gin.H{
			"error": "Remove problem",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "true",
	})
}

func ShowFunctions(c *gin.Context) {
	user, _ := c.Get("userID")
	userID := user.(int)
	db, _ := c.Get("db")
	database := db.(*gorm.DB)
	functions := structs.GetFunctions(userID, database)
	c.JSON(200, gin.H{
		"totalResult": len(functions),
		"result":      functions,
	})
}
