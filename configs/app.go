package configs

type App struct {
	Env     string `mapstructure:"ENV" json:"ENV" yaml:"ENV"`
	Port    string `mapstructure:"PORT" json:"PORT" yaml:"PORT"`
	AppName string `mapstructure:"APP_NAME" json:"APP_NAME" yaml:"APP_NAME"`
	AppUrl  string `mapstructure:"APP_URL" json:"APP_URL" yaml:"APP_URL"`
}
