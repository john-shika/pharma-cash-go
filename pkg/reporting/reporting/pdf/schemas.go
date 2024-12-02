package pdf

import (
	"github.com/shopspring/decimal"
	"time"
)

type FormDataPdf struct {
	Title string    `mapstructure:"title" json:"title"`
	Name  string    `mapstructure:"name" json:"name"`
	Role  string    `mapstructure:"role" json:"role"`
	Date  time.Time `mapstructure:"date" json:"date"`
}

func (FormDataPdf) GetNameType() string {
	return "FormData"
}

type Page1TableColsPdf struct {
	Name   string          `mapstructure:"name" json:"name"`
	Buy    decimal.Decimal `mapstructure:"buy" json:"buy"`
	Margin decimal.Decimal `mapstructure:"margin" json:"margin"`
	Tax    decimal.Decimal `mapstructure:"tax" json:"tax"`
	Sale   decimal.Decimal `mapstructure:"sale" json:"sale"`
	Stock  int             `mapstructure:"stock" json:"stock"`
	Sold   int             `mapstructure:"sold" json:"sold"`
	Date   time.Time       `mapstructure:"date" json:"date"`
}

func (Page1TableColsPdf) GetNameType() string {
	return "Page1TableCols"
}

type Page1TableRowsPdf []Page1TableColsPdf

func (Page1TableRowsPdf) GetNameType() string {
	return "Page1TableRows"
}

type Page2TableColsPdf struct {
	Name     string          `mapstructure:"name" json:"name"`
	Officer  string          `mapstructure:"officer" json:"officer"`
	Shift    string          `mapstructure:"shift" json:"shift"`
	Quantity int             `mapstructure:"quantity" json:"quantity"`
	Subtotal decimal.Decimal `mapstructure:"subtotal" json:"subtotal"`
	Total    decimal.Decimal `mapstructure:"total" json:"total"`
	Income   decimal.Decimal `mapstructure:"income" json:"income"`
	Date     time.Time       `mapstructure:"date" json:"date"`
}

func (Page2TableColsPdf) GetNameType() string {
	return "Page2TableCols"
}

type Page2TableRowsPdf []Page2TableColsPdf

func (Page2TableRowsPdf) GetNameType() string {
	return "Page2TableRows"
}

type Page3TableColsPdf struct {
	Name     string    `mapstructure:"name" json:"name"`
	Brand    string    `mapstructure:"brand" json:"brand"`
	Supplier string    `mapstructure:"supplier" json:"supplier"`
	Stock    int       `mapstructure:"stock" json:"stock"`
	Expires  time.Time `mapstructure:"expires" json:"expires"`
}

func (Page3TableColsPdf) GetNameType() string {
	return "Page3TableCols"
}

type Page3TableRowsPdf []Page3TableColsPdf

func (Page3TableRowsPdf) GetNameType() string {
	return "Page3TableRows"
}
