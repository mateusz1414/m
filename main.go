package main

import (
	"log"
	"matura/auth"
	"matura/math"
	"matura/polish"
	"matura/structs"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	db := dbConnect()
	router.Use(dbMiddleWeare(db))
	authr := router.Group("/auth")
	{
		authr.POST("/login", auth.LoginUser)
		authr.POST("/register", auth.RegisterUser)
	}

	router.Use(authf())
	mathe := router.Group("/math")
	{
		mathe.POST("/", math.AddFunction)
		mathe.GET("/getAll", math.ShowFunctions)
		mathe.DELETE("/:id", math.DeleteFunction)
	}
	pl := router.Group("/polish")
	{
		pl.POST("/", polish.AddEra)
		pl.GET("/:id", polish.ShowEras)
		pl.DELETE("/:id", polish.DeleteEra)
	}
	books := router.Group("/books")
	{
		books.POST("/", polish.AddReading)
		books.DELETE("/:id", polish.RemoveReading)
		books.PUT(":id", polish.ChangeStatus)
	}
	router.Run()
}

func authf() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		token := strings.Split(authorization, "Bearer ")[1]
		userID := structs.Login(token)
		if userID == 0 {
			c.JSON(401, gin.H{
				"erro": "you must authenticate",
			})
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}

func dbConnect() *gorm.DB {
	dsn := "DEHxBqlf4R:VFCCiwIsB5@tcp(remotemysql.com:3306)/DEHxBqlf4R?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func dbMiddleWeare(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
