package model

import (
	"time"

	"gorm.io/gorm"
)

// Dummy example
type Dummy struct {
	gorm.Model
	ID        int       `gorm:"primaryKey" example:"1"`
	Name      string    `example:"desktop chair"`
	Price     float64   `example:"299.99"`
	CreatedAt time.Time `swaggerignore:"true"`
	UpdatedAt time.Time `swaggerignore:"true"`
	DeletedAt time.Time `gorm:"-;default:null" swaggerignore:"true"`
}

func (Dummy) TableName() string {
	return "dummy"
}

func (d *Dummy) QCreateDummy(db *gorm.DB) error {
	d.CreatedAt = time.Now()
	d.UpdatedAt = d.CreatedAt
	err := db.Create(d).Error
	return err
}

func (d *Dummy) QGetDummies(db *gorm.DB, start int, count int) ([]Dummy, error) {
	var dummies []Dummy
	if err := db.Table("dummy").Select("*").Scan(&dummies).Error; err != nil {
		return nil, err
	}

	return dummies, nil
}

func (d *Dummy) QGetDummy(db *gorm.DB) error {
	err := db.Where("id = ?", d.ID).First(&d).Error
	return err
}
