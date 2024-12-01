package hwd

type Config struct {
	Exec string `mapstructure:"exec" json:"exec" yaml:"exec"`
	DB   string `mapstructure:"db" json:"db" yaml:"db"`
}

func (Config) GetNameType() string {
	return "Hwd"
}
