package config

var Config AppConfig

type AppConfig struct {
	App       ServerConfig   `yaml:"app"`
	MysqlDemo MysqlConfig    `yaml:"mysql_demo"`
	Log       LogConfig      `yaml:"log"`
}
