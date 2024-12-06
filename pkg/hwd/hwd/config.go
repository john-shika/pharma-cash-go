package hwd

type Printer struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}

func (Printer) GetNameType() string {
	return "Printer"
}

type Scanner struct {
	Serial string `mapstructure:"serial" json:"serial" yaml:"serial"`
}

func (Scanner) GetNameType() string {
	return "Scanner"
}

type Config struct {
	DB      string  `mapstructure:"db" json:"db" yaml:"db"`
	Exec    string  `mapstructure:"exec" json:"exec" yaml:"exec"`
	Printer Printer `mapstructure:"printer" json:"printer" yaml:"printer"`
	Scanner Scanner `mapstructure:"scanner" json:"scanner" yaml:"scanner"`
}

func (Config) GetNameType() string {
	return "Hwd"
}
