package models

import (
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	"strings"
)

type Product struct {
	models.BaseModel
	Barcode          string          `db:"barcode" gorm:"unique;index;not null;" mapstructure:"barcode" json:"barcode"`
	Brand            string          `db:"brand" gorm:"index;not null;" mapstructure:"brand" json:"brand"`
	ProductName      string          `db:"product_name" gorm:"index;not null;" mapstructure:"product_name" json:"productName"`
	Supplier         string          `db:"supplier" gorm:"index;not null;" mapstructure:"supplier" json:"supplier"`
	Description      string          `db:"description" gorm:"index;not null;" mapstructure:"description" json:"description"`
	Expires          sqlx.DateOnly   `db:"expires" gorm:"index;not null;" mapstructure:"expires" json:"expires"`
	PurchasePrice    decimal.Decimal `db:"purchase_price" gorm:"index;not null;" mapstructure:"purchase_price" json:"purchasePrice"`
	SupplierDiscount float32         `db:"supplier_discount" gorm:"index;not null;" mapstructure:"supplier_discount" json:"supplierDiscount"`
	VAT              float32         `db:"vat" gorm:"index;not null;" mapstructure:"vat" json:"vat"`
	ProfitMargin     float32         `db:"profit_margin" gorm:"index;not null;" mapstructure:"profit_margin" json:"profitMargin"`
	PackageID        uint            `db:"package_id" gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" mapstructure:"package_id" json:"packageId"`
	PackageTotal     float32         `db:"package_total" gorm:"index;not null;" mapstructure:"package_total" json:"packageTotal"`
	UnitID           uint            `db:"unit_id" gorm:"index;not null;" mapstructure:"unit_id" json:"unitId"`
	UnitAmount       float32         `db:"unit_amount" gorm:"index;not null;" mapstructure:"unit_amount" json:"unitAmount"`
	UnitExtra        float32         `db:"unit_extra" gorm:"index;not null;" mapstructure:"unit_extra" json:"unitExtra"`

	Package    Package    `db:"-" gorm:"foreignKey:PackageID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" mapstructure:"package" json:"package"`
	Unit       Unit       `db:"-" gorm:"foreignKey:UnitID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" mapstructure:"unit" json:"unit"`
	Categories []Category `db:"-" gorm:"many2many:product_categories;" mapstructure:"categories" json:"categories"`
}

func (Product) TableName() string {
	return "products"
}

func (p *Product) CreateCategories(DB *gorm.DB) error {
	var err error

	for i, category := range p.Categories {
		nokocore.KeepVoid(i)

		// passing
		if category.ID != 0 {
			continue
		}

		// searching
		var check Category
		tx := DB.Where("category_name = ?", category.CategoryName).Find(&check)
		if err = tx.Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// passing
		if check.ID != 0 {
			p.Categories[i] = check
			continue
		}

		// create new
		category.UUID = nokocore.NewUUID()
		tx = DB.Create(&category)
		if err = tx.Error; err != nil {
			return err
		}

		// check rows affected
		if tx.RowsAffected < 1 {
			return errors.New("no rows affected")
		}

		// object assign
		p.Categories[i] = category
	}

	return nil
}

func (p *Product) ClearCategories(DB *gorm.DB) error {
	var err error

	if p.ID != 0 {
		// pseudo product
		product := Product{
			BaseModel: models.BaseModel{
				ID: p.ID,
			},
		}

		// remove all registered product categories, product categories get empty
		if err = DB.Model(&product).Association("Categories").Clear(); err != nil {
			return err
		}
	}

	return nil
}

func (p *Product) CategoriesAppend(DB *gorm.DB, names ...string) error {
	for i, name := range names {
		nokocore.KeepVoid(i)
		found := false
		for j, category := range p.Categories {
			nokocore.KeepVoid(j)
			if strings.EqualFold(category.CategoryName, name) {
				found = true
				break
			}
		}

		if !found {
			p.Categories = append(p.Categories, Category{CategoryName: name})
		}
	}

	return p.CreateCategories(DB)
}

func (p *Product) BeforeSave(DB *gorm.DB) (err error) {
	nokocore.KeepVoid(DB)

	// create product categories if not exists
	if err = p.CreateCategories(DB); err != nil {
		return err
	}

	if err = p.ClearCategories(DB); err != nil {
		return err
	}

	return nil
}
