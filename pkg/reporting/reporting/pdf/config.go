package pdf

type TemplateConfig struct {
	PageFile   string `mapstructure:"page_file" json:"pageFile" yaml:"page_file"`
	PageLayout string `mapstructure:"page_layout" json:"pageLayout" yaml:"page_layout"`
	PageSize   string `mapstructure:"page_size" json:"pageSize" yaml:"page_size"`
}

func (TemplateConfig) GetNameType() string {
	return "PdfTemplate"
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
	FontFamily string    `mapstructure:"font_family" json:"fontFamily" yaml:"font_family"`
	FontType   string    `mapstructure:"font_type" json:"fontType" yaml:"font_type"`
	FontSize   float64   `mapstructure:"font_size" json:"fontSize" yaml:"font_size"`
}

func (Config) GetNameType() string {
	return "Pdf"
}
