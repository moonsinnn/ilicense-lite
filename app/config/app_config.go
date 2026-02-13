package config

var Config AppConfig

type AppConfig struct {
	App     ServerConfig `yaml:"app"`
	MysqlDB MysqlConfig  `yaml:"mysql_db"`
	Log     LogConfig    `yaml:"log"`
}
