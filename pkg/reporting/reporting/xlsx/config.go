package xlsx

type XlsxTemplateConfig struct {
	SheetFile string `mapstructure:"sheet_file" json:"sheetFile" yaml:"sheet_file"`
	SheetName string `mapstructure:"sheet_name" json:"sheetName" yaml:"sheet_name"`
}

func (XlsxTemplateConfig) GetNameType() string {
	return "XlsxTemplate"
}

type XlsxTemplates []XlsxTemplateConfig

func (XlsxTemplates) GetNameType() string {
	return "XlsxTemplates"
}

type XlsxConfig struct {
	Assets     string        `mapstructure:"assets" json:"assets" yaml:"assets"`
	Templates  XlsxTemplates `mapstructure:"templates" json:"templates" yaml:"templates"`
	OutputDir  string        `mapstructure:"output_dir" json:"outputDir" yaml:"output_dir"`
	OutputName string        `mapstructure:"output_name" json:"outputName" yaml:"output_name"`
}

func (XlsxConfig) GetNameType() string {
	return "Xlsx"
}
