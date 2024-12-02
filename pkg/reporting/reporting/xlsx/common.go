package xlsx

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"nokowebapi/nokocore"
	"time"
)

var NumberXlsxFormat = "0"
var DateXlsxFormat = "[$-en-ID,1]dddd, dd mmmm yyyy;@"
var CurrencyXlsxFormat = "_-[$Rp-en-ID]* #,##0.00_-;-[$Rp-en-ID]* #,##0.00_-;_-[$Rp-en-ID]* \"-\"??_-;_-@_-"

var xlsxStyleCodes = nokocore.NewMapLock[int]()

var xlsxBorderThin = []excelize.Border{
	{
		Type:  "top",
		Color: "#000000",
		Style: 1,
	},
	{
		Type:  "left",
		Color: "#000000",
		Style: 1,
	},
	{
		Type:  "right",
		Color: "#000000",
		Style: 1,
	},
	{
		Type:  "bottom",
		Color: "#000000",
		Style: 1,
	},
}

func StyleCached(name string, fn func() int) int {
	if !xlsxStyleCodes.HasKey(name) {
		if fn == nil {
			panic("fn is nil")
		}

		temp := fn()
		xlsxStyleCodes.Set(name, temp)
		return temp
	}

	return xlsxStyleCodes.Get(name)
}

func SetFormTitleXlsx(file *excelize.File, sheetName string, title string) (err error) {
	style := StyleCached("Arial-16-Center-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   16,
			},
		}))
	})

	start := fmt.Sprintf("D%d", 4)
	end := fmt.Sprintf("G%d", 5)
	if err = file.SetCellValue(sheetName, start, title); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, start, end, style); err != nil {
		return err
	}

	if err = file.MergeCell(sheetName, start, end); err != nil {
		return err
	}

	return nil
}

func SetFormNameXlsx(file *excelize.File, sheetName string, name string) (err error) {
	style := StyleCached("Arial-12-Left-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	start := fmt.Sprintf("I%d", 3)
	end := fmt.Sprintf("J%d", 3)
	if err = file.SetCellValue(sheetName, start, name); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, start, end, style); err != nil {
		return err
	}

	if err = file.MergeCell(sheetName, start, end); err != nil {
		return err
	}

	return nil
}

func SetFormRoleXlsx(file *excelize.File, sheetName string, role string) (err error) {
	style := StyleCached("Arial-12-Left-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	start := fmt.Sprintf("I%d", 4)
	end := fmt.Sprintf("J%d", 4)
	if err = file.SetCellValue(sheetName, start, role); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, start, end, style); err != nil {
		return err
	}

	if err = file.MergeCell(sheetName, start, end); err != nil {
		return err
	}

	return nil
}

func SetFormDateXlsx(file *excelize.File, sheetName string, date time.Time) (err error) {
	style := StyleCached("Arial-12-Left-Center-Date-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &DateXlsxFormat,
		}))
	})

	start := fmt.Sprintf("I%d", 5)
	end := fmt.Sprintf("J%d", 5)
	if err = file.SetCellValue(sheetName, start, date); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, start, end, style); err != nil {
		return err
	}

	if err = file.MergeCell(sheetName, start, end); err != nil {
		return err
	}

	return nil
}

func SetFormDataXlsx(file *excelize.File, sheetName string, formData *FormDataXlsx) (err error) {
	if err = SetFormTitleXlsx(file, sheetName, formData.Title); err != nil {
		return err
	}

	if err = SetFormNameXlsx(file, sheetName, formData.Name); err != nil {
		return err
	}

	if err = SetFormRoleXlsx(file, sheetName, formData.Role); err != nil {
		return err
	}

	if err = SetFormDateXlsx(file, sheetName, formData.Date); err != nil {
		return err
	}

	return nil
}

// Sheet1

func SetSheet1TableNumberXlsx(file *excelize.File, sheetName string, index int, value int) (err error) {
	style := StyleCached("Arial-12-Center-Center-Number-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &NumberXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("B%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableNameXlsx(file *excelize.File, sheetName string, index int, value string) (err error) {
	style := StyleCached("Arial-12-Left-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	index += 10
	cell := fmt.Sprintf("C%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableBuyXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &CurrencyXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("D%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableMarginXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &CurrencyXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("E%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableTaxXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &CurrencyXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("F%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableSaleXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &CurrencyXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("G%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableStockXlsx(file *excelize.File, sheetName string, index int, value int) (err error) {
	style := StyleCached("Arial-12-Right-Center-Number-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &NumberXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("H%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableSoldXlsx(file *excelize.File, sheetName string, index int, value int) (err error) {
	style := StyleCached("Arial-12-Right-Center-Number-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &NumberXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("I%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableDateXlsx(file *excelize.File, sheetName string, index int, value time.Time) (err error) {
	style := StyleCached("Arial-12-Right-Center-Date-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &DateXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("J%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableXlsx(file *excelize.File, sheetName string, index int, start int, data Sheet1TableColsXlsx) (err error) {
	if err = SetSheet1TableNumberXlsx(file, sheetName, index, start); err != nil {
		return err
	}

	if err = SetSheet1TableNameXlsx(file, sheetName, index, data.Name); err != nil {
		return err
	}

	if err = SetSheet1TableBuyXlsx(file, sheetName, index, data.Buy.InexactFloat64()); err != nil {
		return err
	}

	if err = SetSheet1TableMarginXlsx(file, sheetName, index, data.Margin.InexactFloat64()); err != nil {
		return err
	}

	if err = SetSheet1TableTaxXlsx(file, sheetName, index, data.Tax.InexactFloat64()); err != nil {
		return err
	}

	if err = SetSheet1TableSaleXlsx(file, sheetName, index, data.Sale.InexactFloat64()); err != nil {
		return err
	}

	if err = SetSheet1TableStockXlsx(file, sheetName, index, data.Stock); err != nil {
		return err
	}

	if err = SetSheet1TableSoldXlsx(file, sheetName, index, data.Sold); err != nil {
		return err
	}

	if err = SetSheet1TableDateXlsx(file, sheetName, index, data.Date); err != nil {
		return err
	}

	return nil
}

// Sheet3

func SetSheet2TableNumberXlsx(file *excelize.File, sheetName string, index int, value int) (err error) {
	style := StyleCached("Arial-12-Center-Center-Number-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &NumberXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("B%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableNameXlsx(file *excelize.File, sheetName string, index int, value string) (err error) {
	style := StyleCached("Arial-12-Left-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	index += 10
	cell := fmt.Sprintf("C%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableBrandXlsx(file *excelize.File, sheetName string, index int, value string) (err error) {
	style := StyleCached("Arial-12-Left-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	index += 10
	start := fmt.Sprintf("D%d", index)
	end := fmt.Sprintf("E%d", index)
	if err = file.SetCellValue(sheetName, start, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, start, end, style); err != nil {
		return err
	}

	if err = file.MergeCell(sheetName, start, end); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableSupplierXlsx(file *excelize.File, sheetName string, index int, value string) (err error) {
	style := StyleCached("Arial-12-Left-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	index += 10
	start := fmt.Sprintf("F%d", index)
	end := fmt.Sprintf("H%d", index)
	if err = file.SetCellValue(sheetName, start, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, start, end, style); err != nil {
		return err
	}

	if err = file.MergeCell(sheetName, start, end); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableStockXlsx(file *excelize.File, sheetName string, index int, value int) (err error) {
	style := StyleCached("Arial-12-Right-Center-Number-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &NumberXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("I%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableExpiresXlsx(file *excelize.File, sheetName string, index int, value time.Time) (err error) {
	style := StyleCached("Arial-12-Right-Center-Date-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &DateXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("J%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	if err = file.MergeCell(sheetName, cell, cell); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableXlsx(file *excelize.File, sheetName string, index int, start int, data Sheet2TableColsXlsx) (err error) {
	if err = SetSheet2TableNumberXlsx(file, sheetName, index, start); err != nil {
		return err
	}

	if err = SetSheet2TableNameXlsx(file, sheetName, index, data.Name); err != nil {
		return err
	}

	if err = SetSheet2TableBrandXlsx(file, sheetName, index, data.Brand); err != nil {
		return err
	}

	if err = SetSheet2TableSupplierXlsx(file, sheetName, index, data.Supplier); err != nil {
		return err
	}

	if err = SetSheet2TableStockXlsx(file, sheetName, index, data.Stock); err != nil {
		return err
	}

	if err = SetSheet2TableExpiresXlsx(file, sheetName, index, data.Expires); err != nil {
		return err
	}

	return nil
}

// Sheet4

func SetSheet3TableNumberXlsx(file *excelize.File, sheetName string, index int, value int) (err error) {
	style := StyleCached("Arial-12-Center-Center-Number-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &NumberXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("B%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet3TableNameXlsx(file *excelize.File, sheetName string, index int, value string) (err error) {
	style := StyleCached("Arial-12-Left-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	index += 10
	cell := fmt.Sprintf("C%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet3TableOfficerXlsx(file *excelize.File, sheetName string, index int, value string) (err error) {
	style := StyleCached("Arial-12-Left-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	index += 10
	cell := fmt.Sprintf("D%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet3TableShiftXlsx(file *excelize.File, sheetName string, index int, value string) (err error) {
	style := StyleCached("Arial-12-Center-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	index += 10
	cell := fmt.Sprintf("E%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet3TableQuantityXlsx(file *excelize.File, sheetName string, index int, value int) (err error) {
	style := StyleCached("Arial-12-Right-Center-Number-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &NumberXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("F%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet3TableSubtotalXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &CurrencyXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("G%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet3TableTotalXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &CurrencyXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("H%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet3TableIncomeXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &CurrencyXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("I%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet3TableDateXlsx(file *excelize.File, sheetName string, index int, value time.Time) (err error) {
	style := StyleCached("Arial-12-Right-Center-Date-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: xlsxBorderThin,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &DateXlsxFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("J%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet3TableXlsx(file *excelize.File, sheetName string, index int, start int, data Sheet3TableColsXlsx) (err error) {
	if err = SetSheet3TableNumberXlsx(file, sheetName, index, start); err != nil {
		return err
	}

	if err = SetSheet3TableNameXlsx(file, sheetName, index, data.Name); err != nil {
		return err
	}

	if err = SetSheet3TableOfficerXlsx(file, sheetName, index, data.Officer); err != nil {
		return err
	}

	if err = SetSheet3TableShiftXlsx(file, sheetName, index, data.Shift); err != nil {
		return err
	}

	if err = SetSheet3TableQuantityXlsx(file, sheetName, index, data.Quantity); err != nil {
		return err
	}

	if err = SetSheet3TableSubtotalXlsx(file, sheetName, index, data.Subtotal.InexactFloat64()); err != nil {
		return err
	}

	if err = SetSheet3TableTotalXlsx(file, sheetName, index, data.Total.InexactFloat64()); err != nil {
		return err
	}

	if err = SetSheet3TableIncomeXlsx(file, sheetName, index, data.Income.InexactFloat64()); err != nil {
		return err
	}

	if err = SetSheet3TableDateXlsx(file, sheetName, index, data.Date); err != nil {
		return err
	}

	return nil
}
