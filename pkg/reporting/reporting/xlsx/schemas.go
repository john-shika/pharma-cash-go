package xlsx

import (
	"github.com/shopspring/decimal"
	"time"
)

type FormDataXlsx struct {
	Title string    `mapstructure:"title" json:"title"`
	Name  string    `mapstructure:"name" json:"name"`
	Role  string    `mapstructure:"role" json:"role"`
	Date  time.Time `mapstructure:"date" json:"date"`
}

func (FormDataXlsx) GetNameType() string {
	return "FormData"
}

type Sheet1TableColsXlsx struct {
	Name   string          `mapstructure:"name" json:"name"`
	Buy    decimal.Decimal `mapstructure:"buy" json:"buy"`
	Margin decimal.Decimal `mapstructure:"margin" json:"margin"`
	Tax    decimal.Decimal `mapstructure:"tax" json:"tax"`
	Sale   decimal.Decimal `mapstructure:"sale" json:"sale"`
	Stock  int             `mapstructure:"stock" json:"stock"`
	Sold   int             `mapstructure:"sold" json:"sold"`
	Date   time.Time       `mapstructure:"date" json:"date"`
}

func (Sheet1TableColsXlsx) GetNameType() string {
	return "Sheet1TableCols"
}

type Sheet1TableRowsXlsx []Sheet1TableColsXlsx

func (Sheet1TableRowsXlsx) GetNameType() string {
	return "Sheet1TableRows"
}

type Sheet2TableColsXlsx struct {
	Name     string    `mapstructure:"name" json:"name"`
	Brand    string    `mapstructure:"brand" json:"brand"`
	Supplier string    `mapstructure:"supplier" json:"supplier"`
	Stock    int       `mapstructure:"stock" json:"stock"`
	Expires  time.Time `mapstructure:"expires" json:"expires"`
}

func (Sheet2TableColsXlsx) GetNameType() string {
	return "Sheet2TableCols"
}

type Sheet2TableRowsXlsx []Sheet2TableColsXlsx

func (Sheet2TableRowsXlsx) GetNameType() string {
	return "Sheet2TableRows"
}

type Sheet3TableColsXlsx struct {
	Name     string          `mapstructure:"name" json:"name"`
	Officer  string          `mapstructure:"officer" json:"officer"`
	Shift    string          `mapstructure:"shift" json:"shift"`
	Quantity int             `mapstructure:"quantity" json:"quantity"`
	Subtotal decimal.Decimal `mapstructure:"subtotal" json:"subtotal"`
	Total    decimal.Decimal `mapstructure:"total" json:"total"`
	Income   decimal.Decimal `mapstructure:"income" json:"income"`
	Date     time.Time       `mapstructure:"date" json:"date"`
}

func (Sheet3TableColsXlsx) GetNameType() string {
	return "Sheet3TableCols"
}

type Sheet3TableRowsXlsx []Sheet3TableColsXlsx

func (Sheet3TableRowsXlsx) GetNameType() string {
	return "Sheet3TableRows"
}
