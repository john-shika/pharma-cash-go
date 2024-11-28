package models

type ProductCategory struct {
	ProductID  uint     `db:"product_id" gorm:"index" mapstructure:"product_id" json:"productId"`
	CategoryID uint     `db:"category_id" gorm:"index" mapstructure:"category_id" json:"categoryId"`
	Product    Product  `db:"-" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"product" json:"product"`
	Category   Category `db:"-" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"category" json:"category"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}
