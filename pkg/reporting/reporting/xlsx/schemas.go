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
	Name     string          `mapstructure:"name" json:"name"`
	Buy      decimal.Decimal `mapstructure:"buy" json:"buy"`
	Margin   decimal.Decimal `mapstructure:"margin" json:"margin"`
	Tax      decimal.Decimal `mapstructure:"tax" json:"tax"`
	Sale     decimal.Decimal `mapstructure:"sale" json:"sale"`
	StockIn  int             `mapstructure:"stock_in" json:"stockIn"`
	StockOut int             `mapstructure:"stock_out" json:"stockOut"`
	Date     time.Time       `mapstructure:"date" json:"date"`
}

func (Sheet1TableColsXlsx) GetNameType() string {
	return "Sheet1TableCols"
}

type Sheet1TableRowsXlsx []Sheet1TableColsXlsx

func (Sheet1TableRowsXlsx) GetNameType() string {
	return "Sheet1TableRows"
}

type Sheet2TableColsXlsx struct {
	Name           string          `mapstructure:"name" json:"name"`
	OfficerName    string          `mapstructure:"officer_name" json:"officerName"`
	OfficerShift   string          `mapstructure:"officer_shift" json:"officerShift"`
	StockIn        int             `mapstructure:"stock_in" json:"stockIn"`
	StockOut       int             `mapstructure:"stock_out" json:"stockOut"`
	SubtotalBuy    decimal.Decimal `mapstructure:"subtotal_buy" json:"subtotalBuy"`
	SubtotalMargin decimal.Decimal `mapstructure:"subtotal_margin" json:"subtotalMargin"`
	SubtotalTax    decimal.Decimal `mapstructure:"subtotal_tax" json:"subtotalTax"`
	SubtotalSale   decimal.Decimal `mapstructure:"subtotal_sale" json:"subtotalSale"`
	TotalBuy       decimal.Decimal `mapstructure:"total_buy" json:"totalBuy"`
	TotalMargin    decimal.Decimal `mapstructure:"total_margin" json:"totalMargin"`
	TotalTax       decimal.Decimal `mapstructure:"total_tax" json:"totalTax"`
	TotalSale      decimal.Decimal `mapstructure:"total_sale" json:"totalSale"`
	Date           time.Time       `mapstructure:"date" json:"date"`
}

func (Sheet2TableColsXlsx) GetNameType() string {
	return "Sheet2TableCols"
}

type Sheet2TableRowsXlsx []Sheet2TableColsXlsx

func (Sheet2TableRowsXlsx) GetNameType() string {
	return "Sheet2TableRows"
}
