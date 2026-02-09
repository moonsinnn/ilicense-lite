package config

type ServerConfig struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}
