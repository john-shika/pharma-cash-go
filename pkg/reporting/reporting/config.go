package reporting

type Config struct {
	Assets     string `mapstructure:"assets" json:"assets" yaml:"assets"`
	OutputDir  string `mapstructure:"output_dir" json:"outputDir" yaml:"output_dir"`
	FontFamily string `mapstructure:"font_family" json:"fontFamily" yaml:"font_family"`
	FontType   string `mapstructure:"font_type" json:"fontType" yaml:"font_type"`
	FontSize   int    `mapstructure:"font_size" json:"fontSize" yaml:"font_size"`
}

func (c *Config) GetNameType() string {
	return "Reporting"
}
