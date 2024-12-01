package xlsx

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"nokowebapi/nokocore"
	"path/filepath"
	"time"
)

type DocXlsxImpl interface {
	Sheet1Print(sheetName string, data *FormDataXlsx, rows Sheet1TableRowsXlsx) error
	Sheet2Print(sheetName string, data *FormDataXlsx, rows Sheet2TableRowsXlsx) error
	Save() error
}

type DocXlsx struct {
	Config *Config
	Excel  *excelize.File
	Sheet  int
	Date   time.Time
}

func NewDocXlsx(templateId int, config *Config) DocXlsxImpl {
	var err error
	var excel *excelize.File
	nokocore.KeepVoid(err, excel)

	template := config.Templates[templateId]
	sheetFilePath := filepath.Join(config.Assets, template.SheetFile)
	fmt.Println(sheetFilePath)

	if excel, err = excelize.OpenFile(sheetFilePath); err != nil {
		panic(fmt.Errorf("failed to open excel, %w", err))
	}

	return &DocXlsx{
		Config: config,
		Excel:  excel,
		Date:   time.Now(),
		Sheet:  1,
	}
}

func (d *DocXlsx) Sheet1Print(sheetName string, data *FormDataXlsx, rows Sheet1TableRowsXlsx) (err error) {
	if err = SetFormTitleXlsx(d.Excel, sheetName, data.Title); err != nil {
		return err
	}

	if err = SetFormNameXlsx(d.Excel, sheetName, data.Name); err != nil {
		return err
	}

	if err = SetFormRoleXlsx(d.Excel, sheetName, data.Role); err != nil {
		return err
	}

	if err = SetFormDateXlsx(d.Excel, sheetName, data.Date); err != nil {
		return err
	}

	for i, row := range rows {
		nokocore.KeepVoid(i)
		j := i + 1

		if err = SetSheet1TableNumberXlsx(d.Excel, sheetName, i, j); err != nil {
			return err
		}

		if err = SetSheet1TableNameXlsx(d.Excel, sheetName, i, row.Name); err != nil {
			return err
		}

		if err = SetSheet1TableBuyXlsx(d.Excel, sheetName, i, row.Buy.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet1TableMarginXlsx(d.Excel, sheetName, i, row.Margin.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet1TableTaxXlsx(d.Excel, sheetName, i, row.Tax.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet1TableSaleXlsx(d.Excel, sheetName, i, row.Sale.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet1TableStockInXlsx(d.Excel, sheetName, i, row.StockIn); err != nil {
			return err
		}

		if err = SetSheet1TableStockOutXlsx(d.Excel, sheetName, i, row.StockOut); err != nil {
			return err
		}

		if err = SetSheet1TableDateXlsx(d.Excel, sheetName, i, row.Date); err != nil {
			return err
		}

	}

	return nil
}

func (d *DocXlsx) Sheet2Print(sheetName string, data *FormDataXlsx, rows Sheet2TableRowsXlsx) (err error) {
	if err = SetFormTitleXlsx(d.Excel, sheetName, data.Title); err != nil {
		return err
	}

	if err = SetFormNameXlsx(d.Excel, sheetName, data.Name); err != nil {
		return err
	}

	if err = SetFormRoleXlsx(d.Excel, sheetName, data.Role); err != nil {
		return err
	}

	if err = SetFormDateXlsx(d.Excel, sheetName, data.Date); err != nil {
		return err
	}

	for i, row := range rows {
		nokocore.KeepVoid(i)
		j := i + 1

		if err = SetSheet1TableNumberXlsx(d.Excel, sheetName, i, j); err != nil {
			return err
		}

		if err = SetSheet1TableNameXlsx(d.Excel, sheetName, i, row.Name); err != nil {
			return err
		}

		if err = SetSheet2TableOfficerNameXlsx(d.Excel, sheetName, i, row.OfficerName); err != nil {
			return err
		}

		if err = SetSheet2TableOfficerShiftXlsx(d.Excel, sheetName, i, row.OfficerShift); err != nil {
			return err
		}

		if err = SetSheet2TableStockInXlsx(d.Excel, sheetName, i, row.StockIn); err != nil {
			return err
		}

		if err = SetSheet2TableStockOutXlsx(d.Excel, sheetName, i, row.StockOut); err != nil {
			return err
		}

		if err = SetSheet2TableSubtotalBuyXlsx(d.Excel, sheetName, i, row.SubtotalBuy.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableSubtotalMarginXlsx(d.Excel, sheetName, i, row.SubtotalMargin.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableSubtotalTaxXlsx(d.Excel, sheetName, i, row.SubtotalTax.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableSubtotalSaleXlsx(d.Excel, sheetName, i, row.SubtotalSale.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableTotalBuyXlsx(d.Excel, sheetName, i, row.TotalBuy.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableTotalMarginXlsx(d.Excel, sheetName, i, row.TotalMargin.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableTotalTaxXlsx(d.Excel, sheetName, i, row.TotalTax.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableTotalSaleXlsx(d.Excel, sheetName, i, row.TotalSale.InexactFloat64()); err != nil {
			return err
		}

		if err = SetSheet2TableDateXlsx(d.Excel, sheetName, i, row.Date); err != nil {
			return err
		}

	}

	return nil
}

func (d *DocXlsx) Save() (err error) {
	outputFilePath := GetOutputNameForDocXlsx(d.Config, d.Sheet, d.Date)
	if err = d.Excel.SaveAs(outputFilePath); err != nil {
		return err
	}

	return nil
}
