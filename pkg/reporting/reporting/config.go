package reporting

type PdfTemplateConfig struct {
	PageFile   string `mapstructure:"page_file" json:"pageFile" yaml:"page_file"`
	PageLayout string `mapstructure:"page_layout" json:"pageLayout" yaml:"page_layout"`
	PageSize   string `mapstructure:"page_size" json:"pageSize" yaml:"page_size"`
}

func (c *PdfTemplateConfig) GetNameType() string {
	return "PdfTemplate"
}

type PdfTemplates []PdfTemplateConfig

func (t *PdfTemplates) GetNameType() string {
	return "PdfTemplates"
}

type PdfConfig struct {
	Assets     string       `mapstructure:"assets" json:"assets" yaml:"assets"`
	Templates  PdfTemplates `mapstructure:"templates" json:"templates" yaml:"templates"`
	OutputDir  string       `mapstructure:"output_dir" json:"outputDir" yaml:"output_dir"`
	OutputName string       `mapstructure:"output_name" json:"outputName" yaml:"output_name"`
	FontFamily string       `mapstructure:"font_family" json:"fontFamily" yaml:"font_family"`
	FontType   string       `mapstructure:"font_type" json:"fontType" yaml:"font_type"`
	FontSize   int          `mapstructure:"font_size" json:"fontSize" yaml:"font_size"`
}

func (c *PdfConfig) GetNameType() string {
	return "Pdf"
}

type XlsxTemplateConfig struct {
	SheetFile string `mapstructure:"sheet_file" json:"sheetFile" yaml:"sheet_file"`
	SheetName string `mapstructure:"sheet_name" json:"sheetName" yaml:"sheet_name"`
}

func (c *XlsxTemplateConfig) GetNameType() string {
	return "XlsxTemplate"
}

type XlsxTemplates []XlsxTemplateConfig

func (t *XlsxTemplates) GetNameType() string {
	return "XlsxTemplates"
}

type XlsxConfig struct {
	Assets     string        `mapstructure:"assets" json:"assets" yaml:"assets"`
	Templates  XlsxTemplates `mapstructure:"templates" json:"templates" yaml:"templates"`
	OutputDir  string        `mapstructure:"output_dir" json:"outputDir" yaml:"output_dir"`
	OutputName string        `mapstructure:"output_name" json:"outputName" yaml:"output_name"`
}

func (c *XlsxConfig) GetNameType() string {
	return "Xlsx"
}

type Config struct {
	Pdf  PdfConfig  `mapstructure:"pdf" json:"pdf" yaml:"pdf"`
	Xlsx XlsxConfig `mapstructure:"xlsx" json:"xlsx" yaml:"xlsx"`
}

func (c *Config) GetNameType() string {
	return "Reporting"
}
