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
	Sheet3Print(sheetName string, data *FormDataXlsx, rows Sheet3TableRowsXlsx) error
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

func (d *DocXlsx) Sheet1Print(sheetName string, formData *FormDataXlsx, rows Sheet1TableRowsXlsx) (err error) {
	if err = SetFormDataXlsx(d.Excel, sheetName, formData); err != nil {
		return err
	}

	for i, row := range rows {
		nokocore.KeepVoid(i)
		j := i + 1

		if err = SetSheet1TableXlsx(d.Excel, sheetName, i, j, row); err != nil {
			return err
		}
	}

	return nil
}

func (d *DocXlsx) Sheet2Print(sheetName string, formData *FormDataXlsx, rows Sheet2TableRowsXlsx) (err error) {
	if err = SetFormDataXlsx(d.Excel, sheetName, formData); err != nil {
		return err
	}

	for i, row := range rows {
		nokocore.KeepVoid(i)
		j := i + 1

		if err = SetSheet2TableXlsx(d.Excel, sheetName, i, j, row); err != nil {
			return err
		}
	}

	return nil
}

func (d *DocXlsx) Sheet3Print(sheetName string, formData *FormDataXlsx, rows Sheet3TableRowsXlsx) (err error) {
	if err = SetFormDataXlsx(d.Excel, sheetName, formData); err != nil {
		return err
	}

	for i, row := range rows {
		nokocore.KeepVoid(i)
		j := i + 1

		if err = SetSheet3TableXlsx(d.Excel, sheetName, i, j, row); err != nil {
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
