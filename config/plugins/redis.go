package plugins

// Redis redis配置
type Redis struct {
	Enable   bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
