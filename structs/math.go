package structs

import (
	"github.com/jinzhu/gorm"
)

type Function struct {
	FunctionID int    `gorm:"column:FunctionID"`
	UserID     int    `gorm:"column:UserID"`
	Value      string `gorm:"column:Value" json:"value"`
}

func FindFunction(userID int, functionID int, db *gorm.DB) (function Function) {
	db.Debug().Table("Math").Where("FunctionID=? AND UserID=?", functionID, userID).First(&function)
	return function
}

func GetFunctions(userID int, db *gorm.DB) (functions []Function) {
	db.Debug().Table("Math").Where("UserID=?", userID).Find(&functions)
	return functions
}

func (function *Function) AddFunction(db *gorm.DB) error {
	db.Debug().Table("Math").Create(&function)
	return db.Error
}

func (function *Function) RemoveFunction(db *gorm.DB) error {
	result := db.Debug().Table("Math").Where("FunctionID=?", function.FunctionID).Delete(&function)
	return result.Error
}
