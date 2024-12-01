package xlsx

type TemplateConfig struct {
	SheetFile string `mapstructure:"sheet_file" json:"sheetFile" yaml:"sheet_file"`
	SheetName string `mapstructure:"sheet_name" json:"sheetName" yaml:"sheet_name"`
}

func (TemplateConfig) GetNameType() string {
	return "XlsxTemplate"
}

type Templates []TemplateConfig

func (Templates) GetNameType() string {
	return "Templates"
}

type Config struct {
	Assets     string    `mapstructure:"assets" json:"assets" yaml:"assets"`
	Templates  Templates `mapstructure:"templates" json:"templates" yaml:"templates"`
	OutputDir  string    `mapstructure:"output_dir" json:"outputDir" yaml:"output_dir"`
	OutputName string    `mapstructure:"output_name" json:"outputName" yaml:"output_name"`
}

func (Config) GetNameType() string {
	return "Xlsx"
}
