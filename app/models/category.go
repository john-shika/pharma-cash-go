package models

import "nokowebapi/apis/models"

type Category struct {
	models.BaseModel
	CategoryName string `db:"category_name" gorm:"unique;index;not null;" mapstructure:"category_name" json:"categoryName"`
}

func (Category) TableName() string {
	return "categories"
}
