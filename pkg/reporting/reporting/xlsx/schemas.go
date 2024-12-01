package xlsx

import (
	"github.com/shopspring/decimal"
	"time"
)

type FormDataXlsx struct {
	Title string    `json:"title"`
	Name  string    `json:"name"`
	Role  string    `json:"role"`
	Date  time.Time `json:"date"`
}

func (FormDataXlsx) GetNameType() string {
	return "FormData"
}

type Sheet1TableColsXlsx struct {
	Name     string          `json:"name"`
	Buy      decimal.Decimal `json:"buy"`
	Margin   decimal.Decimal `json:"margin"`
	Tax      decimal.Decimal `json:"tax"`
	Sale     decimal.Decimal `json:"sale"`
	StockIn  int             `json:"stock_in"`
	StockOut int             `json:"stock_out"`
	Date     time.Time       `json:"date"`
}

func (Sheet1TableColsXlsx) GetNameType() string {
	return "Sheet1TableCols"
}

type Sheet1TableRowsXlsx []Sheet1TableColsXlsx

func (Sheet1TableRowsXlsx) GetNameType() string {
	return "Sheet1TableRows"
}

type Sheet2TableColsXlsx struct {
	Name           string          `json:"name"`
	OfficerName    string          `json:"officer_name"`
	OfficerShift   string          `json:"officer_shift"`
	StockIn        int             `json:"stock_in"`
	StockOut       int             `json:"stock_out"`
	SubtotalBuy    decimal.Decimal `json:"subtotal_buy"`
	SubtotalMargin decimal.Decimal `json:"subtotal_margin"`
	SubtotalTax    decimal.Decimal `json:"subtotal_tax"`
	SubtotalSale   decimal.Decimal `json:"subtotal_sale"`
	TotalBuy       decimal.Decimal `json:"total_buy"`
	TotalMargin    decimal.Decimal `json:"total_margin"`
	TotalTax       decimal.Decimal `json:"total_tax"`
	TotalSale      decimal.Decimal `json:"total_sale"`
	Date           time.Time       `json:"date"`
}

func (Sheet2TableColsXlsx) GetNameType() string {
	return "Sheet2TableCols"
}

type Sheet2TableRowsXlsx []Sheet2TableColsXlsx

func (Sheet2TableRowsXlsx) GetNameType() string {
	return "Sheet2TableRows"
}
