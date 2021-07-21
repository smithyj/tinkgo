package logx

type Config struct {
	Name       string `yaml:"Name"`
	Debug      bool   `yaml:"Debug"`
	Level      string `yaml:"Level"`
	Path       string `yaml:"Path"`
	Compress   bool   `yaml:"Compress"`
	MaxSize    int    `yaml:"MaxSize"`
	MaxAge     int    `yaml:"MaxAge"`
	MaxBackups int    `yaml:"MaxBackups"`
}
