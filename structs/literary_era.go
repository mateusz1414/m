package structs

import (
	"github.com/jinzhu/gorm"
)

type Reading struct {
	ReadingID   int    `gorm:"column:ReadingID"`
	EraID       int    `gorm:"column:EraID" json:"eraID"`
	ReadingName string `gorm:"column:ReadingName" json:"readingName"`
	Author      string `gorm:"column:Author" json:"author"`
	Readed      bool   `gorm:"column:Readed" json:"readed"`
}

type LiteraryEra struct {
	EraID   int       `gorm:"column:EraID"`
	UserID  int       `gorm:"column:UserID"`
	EraName string    `gorm:"column:EraName" json:"name"`
	Years   string    `gorm:"column:Years" json:"years"`
	Reading []Reading `gorm:"foreignkey:EraID;association_foreignkey:EraID"`
}

func GetAll(userID int, db *gorm.DB) (literaryEras []LiteraryEra) {
	db.SingularTable(true)
	db.Debug().Table("LiteraryEra").Where("UserID=?", userID).Preload("Reading").Find(&literaryEras)
	return literaryEras
}

func GetEra(userID int, eraID int, db *gorm.DB) (literaryEra LiteraryEra) {
	db.SingularTable(true)
	db.Debug().Table("LiteraryEra").Where("UserID=? AND EraID=?", userID, eraID).Preload("Reading").First(&literaryEra)
	return literaryEra
}

func AddGroup(data LiteraryEra, db *gorm.DB) error {
	err := db.Debug().Table("LiteraryEra").Create(&data).Error
	return err
}

func RemoveEra(userID int, eraID int, db *gorm.DB) error {
	err := db.Debug().Table("LiteraryEra").Where("UserID=? AND EraID=?", userID, eraID).Delete(&LiteraryEra{}).Error
	if err == nil {
		db.Debug().Table("reading").Where("EraID=?", eraID).Delete(&Reading{})
	}
	return err
}

func AddReading(data Reading, db *gorm.DB) error {
	err := db.Debug().Table("reading").Create(&data).Error
	return err
}

func RemoveReading(readingID int, db *gorm.DB) error {
	err := db.Debug().Table("reading").Where("ReadingID=?", readingID).Delete(&Reading{}).Error
	return err
}

func ChangeReadedStatus(reading Reading, db *gorm.DB) error {
	err := db.Debug().Table("reading").Select("Readed").Where("ReadingID=?", reading.ReadingID).Update(reading).Error
	return err
}
