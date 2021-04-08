package polish

import (
	"matura/structs"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ShowEras(c *gin.Context) {
	user, _ := c.Get("userID")
	userID := user.(int)
	db, _ := c.Get("db")
	database := db.(*gorm.DB)
	if id := c.Param("id"); id != "getAll" {
		OneEra(userID, database, id, c)
		return
	}
	result := structs.GetAll(userID, database)
	c.JSON(200, result)
}

func OneEra(userID int, database *gorm.DB, eraID string, c *gin.Context) {
	eraIntID, _ := strconv.Atoi(eraID)
	result := structs.GetEra(userID, eraIntID, database)
	c.JSON(200, result)
}

func AddEra(c *gin.Context) {
	era := structs.LiteraryEra{}
	c.BindJSON(&era)
	user, _ := c.Get("userID")
	era.UserID = user.(int)
	db, _ := c.Get("db")
	database := db.(*gorm.DB)
	err := structs.AddGroup(era, database)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Add era problem",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "true",
	})

}

func DeleteEra(c *gin.Context) {
	eraID := c.Param("id")
	eraIDInt, _ := strconv.Atoi(eraID)
	user, _ := c.Get("userID")
	userID := user.(int)
	db, _ := c.Get("db")
	database := db.(*gorm.DB)
	err := structs.RemoveEra(userID, eraIDInt, database)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Delete era problem",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "true",
	})

}

func AddReading(c *gin.Context) {
	reading := structs.Reading{}
	c.BindJSON(&reading)
	var a = false
	reading.Readed = &a
	db, _ := c.Get("db")
	database := db.(*gorm.DB)
	err := structs.AddReading(reading, database)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Add reading problem",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "true",
	})
}

func RemoveReading(c *gin.Context) {
	readingID := c.Param("id")
	readingIDInt, _ := strconv.Atoi(readingID)
	db, _ := c.Get("db")
	database := db.(*gorm.DB)
	err := structs.RemoveReading(readingIDInt, database)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Remove reading problem",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "true",
	})
}

func ChangeStatus(c *gin.Context) {
	reading := structs.Reading{}
	readingID := c.Param("id")
	c.BindJSON(&reading)
	reading.ReadingID, _ = strconv.Atoi(readingID)
	db, _ := c.Get("db")
	database := db.(*gorm.DB)
	err := structs.ChangeReadedStatus(reading, database)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Remove reading problem",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "true",
	})
}
