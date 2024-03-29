package entity

import "github.com/google/uuid"

type Bmi struct {
	Id          uuid.UUID `gorm:"primaryKey;column:bmi_id;type:varchar(36)"`
	Kg          float64   `gorm:"column:kg"`
	M           float64   `gorm:"column:m"`
	Bmi         float64   `gorm:"column:bmi"`
	Description string    `gorm:"column:description;type:varchar(100)"`
}

func (Bmi) TableName() string {
	return "tb_bmi"
}
