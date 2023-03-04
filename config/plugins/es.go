package plugins

type ES struct {
	Enable bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	Url    string `mapstructure:"url" json:"url" yaml:"url"`
	Index  string `mapstructure:"index" json:"index" yaml:"index"`
}
