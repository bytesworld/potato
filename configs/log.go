package configs

type Log struct {
	LogLevel   string `mapstructure:"LEVEL" json:"LEVEL" yaml:"LEVEL"`
	Format     string `mapstructure:"FORMAT" json:"FORMAT" yaml:"FORMAT"`
	ShowLine   bool   `mapstructure:"SHOW_LINE" json:"SHOW_LINE" yaml:"SHOW_LINE"`
	Path       string `mapstructure:"PATH" json:"PATH" yaml:"PATH"`
	FileName   string `mapstructure:"FILE_NAME" json:"FILE_NAME" yaml:"FILE_NAME"`
	MaxBackups string `mapstructure:"MAX_BACKUPS" json:"MAX_BACKUPS" yaml:"MAX_BACKUPS"`
	MaxSize    string `mapstructure:"MAX_SIZE" json:"MAX_SIZE" yaml:"MAX_SIZE"`
	MaxAge     string `mapstructure:"MAX_AGE" json:"MAX_AGE" yaml:"MAX_AGE"`
}
