package reporting

import (
	"reporting/reporting/pdf"
	"reporting/reporting/xlsx"
)

type Config struct {
	Pdf  pdf.Config  `mapstructure:"pdf" json:"pdf" yaml:"pdf"`
	Xlsx xlsx.Config `mapstructure:"xlsx" json:"xlsx" yaml:"xlsx"`
}

func (Config) GetNameType() string {
	return "Reporting"
}
