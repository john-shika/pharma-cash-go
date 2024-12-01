package pdf

import (
	"github.com/shopspring/decimal"
	"time"
)

type FormDataPdf struct {
	Name string    `mapstructure:"name" json:"name"`
	Role string    `mapstructure:"role" json:"role"`
	Date time.Time `mapstructure:"date" json:"date"`
}

func (FormDataPdf) GetNameType() string {
	return "FormData"
}

type Page1TableColsPdf struct {
	Name     string          `mapstructure:"name" json:"name"`
	Buy      decimal.Decimal `mapstructure:"buy" json:"buy"`
	Margin   decimal.Decimal `mapstructure:"margin" json:"margin"`
	Tax      decimal.Decimal `mapstructure:"tax" json:"tax"`
	Sale     decimal.Decimal `mapstructure:"sale" json:"sale"`
	StockIn  int             `mapstructure:"stock_in" json:"stockIn"`
	StockOut int             `mapstructure:"stock_out" json:"stockOut"`
	Date     time.Time       `mapstructure:"date" json:"date"`
}

func (Page1TableColsPdf) GetNameType() string {
	return "Page1TableCols"
}

type Page1TableRowsPdf []Page1TableColsPdf

func (Page1TableRowsPdf) GetNameType() string {
	return "Page1TableRows"
}
