package configs

type Db struct {
	Engine     string `mapstructure:"ENGINE" json:"ENGINE" yaml:"ENGINE"`
	Name    string `mapstructure:"NAME" json:"NAME" yaml:"NAME"`
}