package bootstrap

import (
	"fmt"
	"log"
	"os"

	"ilicense-lite/config"

	"gopkg.in/yaml.v3"
)

func InitConfig(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	err = yaml.Unmarshal(data, &config.Config)
	if err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}
	applyEnvOverrides()
}

func applyEnvOverrides() {
	if v := os.Getenv("APP_PORT"); v != "" {
		var port int
		if _, err := fmt.Sscanf(v, "%d", &port); err == nil && port > 0 {
			config.Config.App.Port = port
		}
	}
	if v := os.Getenv("JWT_SECRET"); v != "" {
		config.Config.App.JWTSecret = v
	}
	if v := os.Getenv("MYSQL_PASSWORD"); v != "" {
		config.Config.MysqlDB.DataSource.Password = v
	}
	if v := os.Getenv("MYSQL_USERNAME"); v != "" {
		config.Config.MysqlDB.DataSource.UserName = v
	}
	if v := os.Getenv("MYSQL_ADDRESS"); v != "" {
		config.Config.MysqlDB.DataSource.Address = v
	}
	if v := os.Getenv("MYSQL_DB_NAME"); v != "" {
		config.Config.MysqlDB.DataSource.DBName = v
	}
}
