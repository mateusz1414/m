package structs

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	UserID   int    `gorm:"column:UserID"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

func (user *User) Find(db *gorm.DB) error {
	user.UserID = 0
	db.Debug().Table("Users").Where("email=?", user.Email).First(&user)
	if user.UserID == 0 {
		return fmt.Errorf("User not found")
	}
	return nil
}

func (user *User) Exist(db *gorm.DB) bool {
	var count int = 0
	db.Debug().Table("Users").Where("email=?", user.Email).Count(&count)
	if count != 0 {
		return true
	}
	return false
}

func (user *User) Create(db *gorm.DB) error {
	response := db.Debug().Table("Users").Select("email", "password").Create(&user)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (user *User) GenerateToken() (string, error) {
	claims := jwt.MapClaims{
		"userID": user.UserID,
		"email":  user.Email,
		"exp":    time.Now().Unix() + 3600,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	authToken, err := token.SignedString([]byte("dsaggffd"))
	return authToken, err
}

func Login(token string) int {
	tok, err := jwt.Parse(token, func(tok *jwt.Token) (interface{}, error) {
		return []byte("dsaggffd"), nil
	})
	if err != nil {
		return 0
	}
	if claims, ok := tok.Claims.(jwt.MapClaims); ok && int64(claims["exp"].(float64)) > time.Now().Unix() {
		return int(claims["userID"].(float64))

	}
	return 0
}
