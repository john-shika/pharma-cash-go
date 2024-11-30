package xlsx

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"nokowebapi/nokocore"
	"time"
)

var xlsxNumberFormat = "0"
var xlsxDateFormat = "[$-en-ID,1]dddd, dd mmmm yyyy;@"
var xlsxCurrencyFormat = "_-[$Rp-en-ID]* #,##0.00_-;-[$Rp-en-ID]* #,##0.00_-;_-[$Rp-en-ID]* \"-\"??_-;_-@_-"

var styleCodes = nokocore.NewMapLock[int]()

var border = []excelize.Border{
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
	if !styleCodes.HasKey(name) {
		if fn == nil {
			panic("fn is nil")
		}

		temp := fn()
		styleCodes.Set(name, temp)
		return temp
	}

	return styleCodes.Get(name)
}

func SetFormTitleXlsx(file *excelize.File, sheet string, title string) (err error) {
	style := StyleCached("Arial-16-Center-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   16,
			},
		}))
	})

	if err = file.SetCellValue(sheet, "D4", title); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, "D4", "G5", style); err != nil {
		return err
	}

	if err = file.MergeCell(sheet, "D4", "G5"); err != nil {
		return err
	}

	return nil
}

func SetFormNameXlsx(file *excelize.File, sheet string, name string) (err error) {
	style := StyleCached("Arial-12-Left-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	if err = file.SetCellValue(sheet, "I3", name); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, "I3", "J3", style); err != nil {
		return err
	}

	if err = file.MergeCell(sheet, "I3", "J3"); err != nil {
		return err
	}

	return nil
}

func SetFormRoleXlsx(file *excelize.File, sheet string, role string) (err error) {
	style := StyleCached("Arial-12-Left-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	if err = file.SetCellValue(sheet, "I4", role); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, "I4", "J4", style); err != nil {
		return err
	}

	if err = file.MergeCell(sheet, "I4", "J4"); err != nil {
		return err
	}

	return nil
}

func SetFormDateXlsx(file *excelize.File, sheet string, date time.Time) (err error) {
	style := StyleCached("Arial-12-Left-Center-Date-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxDateFormat,
		}))
	})

	if err = file.SetCellValue(sheet, "I5", date); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, "I5", "J5", style); err != nil {
		return err
	}

	if err = file.MergeCell(sheet, "I5", "J5"); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableNumberXlsx(file *excelize.File, sheet string, index int, value int) (err error) {
	style := StyleCached("Arial-12-Center-Center-Number-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxNumberFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("B%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableNameXlsx(file *excelize.File, sheet string, index int, value string) (err error) {
	style := StyleCached("Arial-12-Left-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	index += 10
	cell := fmt.Sprintf("C%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableBuyXlsx(file *excelize.File, sheet string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxCurrencyFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("D%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableMarginXlsx(file *excelize.File, sheet string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxCurrencyFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("E%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableTaxXlsx(file *excelize.File, sheet string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxCurrencyFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("F%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableSaleXlsx(file *excelize.File, sheet string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxCurrencyFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("G%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableStockInXlsx(file *excelize.File, sheet string, index int, value int) (err error) {
	style := StyleCached("Arial-12-Right-Center-Number-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxNumberFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("H%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableStockOutXlsx(file *excelize.File, sheet string, index int, value int) (err error) {
	style := StyleCached("Arial-12-Right-Center-Number-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxNumberFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("I%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet1TableDateXlsx(file *excelize.File, sheet string, index int, value time.Time) (err error) {
	style := StyleCached("Arial-12-Right-Center-Date-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxDateFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("J%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableOfficerNameXlsx(file *excelize.File, sheet string, index int, value string) (err error) {
	style := StyleCached("Arial-12-Left-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "left",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	index += 10
	cell := fmt.Sprintf("D%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableOfficerShiftXlsx(file *excelize.File, sheet string, index int, value string) (err error) {
	style := StyleCached("Arial-12-Center-Center", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
		}))
	})

	index += 10
	cell := fmt.Sprintf("E%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableStockInXlsx(file *excelize.File, sheet string, index int, value int) (err error) {
	style := StyleCached("Arial-12-Right-Center-Number-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxNumberFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("F%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableStockOutXlsx(file *excelize.File, sheet string, index int, value int) (err error) {
	style := StyleCached("Arial-12-Right-Center-Number-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxNumberFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("G%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableSubtotalBuyXlsx(file *excelize.File, sheet string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxCurrencyFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("H%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableSubtotalMarginXlsx(file *excelize.File, sheet string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxCurrencyFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("I%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableSubtotalTaxXlsx(file *excelize.File, sheet string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxCurrencyFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("J%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableSubtotalSaleXlsx(file *excelize.File, sheet string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxCurrencyFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("K%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableTotalBuyXlsx(file *excelize.File, sheet string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxCurrencyFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("L%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableTotalMarginXlsx(file *excelize.File, sheet string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxCurrencyFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("M%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableTotalTaxXlsx(file *excelize.File, sheet string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxCurrencyFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("N%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableTotalSaleXlsx(file *excelize.File, sheet string, index int, value float64) (err error) {
	style := StyleCached("Arial-12-Right-Center-Currency-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxCurrencyFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("O%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}

func SetSheet2TableDateXlsx(file *excelize.File, sheet string, index int, value time.Time) (err error) {
	style := StyleCached("Arial-12-Right-Center-Date-Format", func() int {
		return nokocore.Unwrap(file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
			Border: border,
			Font: &excelize.Font{
				Family: "Arial",
				Size:   12,
			},
			CustomNumFmt: &xlsxDateFormat,
		}))
	})

	index += 10
	cell := fmt.Sprintf("P%d", index)
	if err = file.SetCellValue(sheet, cell, value); err != nil {
		return err
	}

	if err = file.SetCellStyle(sheet, cell, cell, style); err != nil {
		return err
	}

	return nil
}
