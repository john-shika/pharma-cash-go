package xlsx

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"nokowebapi/nokocore"
	"path/filepath"
	"time"
)

type DocXlsxImpl interface {
	Sheet1Print(sheetName string, formDataXlsx *FormDataXlsx, rows Sheet1TableRowsXlsx) error
	Sheet2Print(sheetName string, formDataXlsx *FormDataXlsx, rows Sheet2TableRowsXlsx) error
	Save() error
}

type DocXlsx struct {
	Config *XlsxConfig
	File   *excelize.File
	Date   time.Time
	Sheet  int
}

func NewDocXlsx(config *XlsxConfig) DocXlsxImpl {
	var err error
	var file *excelize.File
	nokocore.KeepVoid(err, file)

	index := 0
	template := config.Templates[index]
	sheetFilePath := filepath.Join(config.Assets, template.SheetFile)
	fmt.Println(sheetFilePath)

	if file, err = excelize.OpenFile(sheetFilePath); err != nil {
		panic(fmt.Errorf("failed to open file, %w", err))
	}

	return &DocXlsx{
		Config: config,
		File:   file,
		Date:   time.Now(),
		Sheet:  index,
	}
}

func (d *DocXlsx) Sheet1Print(sheetName string, formDataXlsx *FormDataXlsx, rows Sheet1TableRowsXlsx) (err error) {
	if err = SetFormTitleXlsx(d.File, sheetName, "Pharma Cash App"); err != nil {
		return err
	}

	if err = SetFormNameXlsx(d.File, sheetName, "John, Doe"); err != nil {
		return err
	}

	if err = SetFormRoleXlsx(d.File, sheetName, "Administrator"); err != nil {
		return err
	}

	if err = SetFormDateXlsx(d.File, sheetName, d.Date); err != nil {
		return err
	}

	for i, row := range rows {
		nokocore.KeepVoid(i)
		j := i + 1

		if err = SetSheet1TableNumberXlsx(d.File, sheetName, i, j); err != nil {
			return err
		}

		if err = SetSheet1TableNameXlsx(d.File, sheetName, i, row.Name); err != nil {
			return err
		}

		if err = SetSheet1TableBuyXlsx(d.File, sheetName, i, row.Buy.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet1TableMarginXlsx(d.File, sheetName, i, row.Margin.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet1TableTaxXlsx(d.File, sheetName, i, row.Tax.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet1TableSaleXlsx(d.File, sheetName, i, row.Sale.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet1TableStockInXlsx(d.File, sheetName, i, row.StockIn); err != nil {
			return err
		}

		if err = SetSheet1TableStockOutXlsx(d.File, sheetName, i, row.StockOut); err != nil {
			return err
		}

		if err = SetSheet1TableDateXlsx(d.File, sheetName, i, row.Date); err != nil {
			return err
		}

	}

	return nil
}

func (d *DocXlsx) Sheet2Print(sheetName string, formDataXlsx *FormDataXlsx, rows Sheet2TableRowsXlsx) (err error) {
	if err = SetFormTitleXlsx(d.File, sheetName, "Pharma Cash App"); err != nil {
		return err
	}

	if err = SetFormNameXlsx(d.File, sheetName, "John, Doe"); err != nil {
		return err
	}

	if err = SetFormRoleXlsx(d.File, sheetName, "Administrator"); err != nil {
		return err
	}

	if err = SetFormDateXlsx(d.File, sheetName, d.Date); err != nil {
		return err
	}

	for i, row := range rows {
		nokocore.KeepVoid(i)
		j := i + 1

		if err = SetSheet1TableNumberXlsx(d.File, sheetName, i, j); err != nil {
			return err
		}

		if err = SetSheet1TableNameXlsx(d.File, sheetName, i, row.Name); err != nil {
			return err
		}

		if err = SetSheet2TableOfficerNameXlsx(d.File, sheetName, i, row.OfficerName); err != nil {
			return err
		}

		if err = SetSheet2TableOfficerShiftXlsx(d.File, sheetName, i, row.OfficerShift); err != nil {
			return err
		}

		if err = SetSheet2TableStockInXlsx(d.File, sheetName, i, row.StockIn); err != nil {
			return err
		}

		if err = SetSheet2TableStockOutXlsx(d.File, sheetName, i, row.StockOut); err != nil {
			return err
		}

		if err = SetSheet2TableSubtotalBuyXlsx(d.File, sheetName, i, row.SubtotalBuy.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableSubtotalMarginXlsx(d.File, sheetName, i, row.SubtotalMargin.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableSubtotalTaxXlsx(d.File, sheetName, i, row.SubtotalTax.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableSubtotalSaleXlsx(d.File, sheetName, i, row.SubtotalSale.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableTotalBuyXlsx(d.File, sheetName, i, row.TotalBuy.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableTotalMarginXlsx(d.File, sheetName, i, row.TotalMargin.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableTotalTaxXlsx(d.File, sheetName, i, row.TotalTax.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableTotalSaleXlsx(d.File, sheetName, i, row.TotalSale.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableDateXlsx(d.File, sheetName, i, row.Date); err != nil {
			return err
		}

	}

	return nil
}

func (d *DocXlsx) Save() (err error) {
	outputFilePath := GetOutputNameForDocXlsx(d.Config, d.Sheet, d.Date)
	if err = d.File.SaveAs(outputFilePath); err != nil {
		return err
	}

	return nil
}
