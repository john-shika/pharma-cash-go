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
		Style: 2,
	},
	{
		Type:  "left",
		Color: "#000000",
		Style: 2,
	},
	{
		Type:  "right",
		Color: "#000000",
		Style: 2,
	},
	{
		Type:  "bottom",
		Color: "#000000",
		Style: 2,
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

	if err = file.SetCellValue(sheetName, "D4", title); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, "D4", "G5", style); err != nil {
		return err
	}

	if err = file.MergeCell(sheetName, "D4", "G5"); err != nil {
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

	if err = file.SetCellValue(sheetName, "I3", name); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, "I3", "J3", style); err != nil {
		return err
	}

	if err = file.MergeCell(sheetName, "I3", "J3"); err != nil {
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

	if err = file.SetCellValue(sheetName, "I4", role); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, "I4", "J4", style); err != nil {
		return err
	}

	if err = file.MergeCell(sheetName, "I4", "J4"); err != nil {
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

	if err = file.SetCellValue(sheetName, "I5", date); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, "I5", "J5", style); err != nil {
		return err
	}

	if err = file.MergeCell(sheetName, "I5", "J5"); err != nil {
		return err
	}

	return nil
}

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

func SetSheet1TableStockInXlsx(file *excelize.File, sheetName string, index int, value int) (err error) {
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

func SetSheet1TableStockOutXlsx(file *excelize.File, sheetName string, index int, value int) (err error) {
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

func SetSheet2TableOfficerNameXlsx(file *excelize.File, sheetName string, index int, value string) (err error) {
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
	cell := fmt.Sprintf("D%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableOfficerShiftXlsx(file *excelize.File, sheetName string, index int, value string) (err error) {
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

func SetSheet2TableStockInXlsx(file *excelize.File, sheetName string, index int, value int) (err error) {
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

func SetSheet2TableStockOutXlsx(file *excelize.File, sheetName string, index int, value int) (err error) {
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
	cell := fmt.Sprintf("G%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableSubtotalBuyXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
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

func SetSheet2TableSubtotalMarginXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
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

func SetSheet2TableSubtotalTaxXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
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
	cell := fmt.Sprintf("J%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableSubtotalSaleXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
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
	cell := fmt.Sprintf("K%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableTotalBuyXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
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
	cell := fmt.Sprintf("L%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableTotalMarginXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
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
	cell := fmt.Sprintf("M%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableTotalTaxXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
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
	cell := fmt.Sprintf("N%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableTotalSaleXlsx(file *excelize.File, sheetName string, index int, value float64) (err error) {
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
	cell := fmt.Sprintf("O%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableDateXlsx(file *excelize.File, sheetName string, index int, value time.Time) (err error) {
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
	cell := fmt.Sprintf("P%d", index)
	if err = file.SetCellValue(sheetName, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheetName, cell, cell, style); err != nil {
		return err
	}

	return nil
}
