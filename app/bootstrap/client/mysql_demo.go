package client

import (
	"ilicense-lite/config"
	"ilicense-lite/library/mysql"

	"gorm.io/gorm"
)

var MysqlDemo *gorm.DB

func InitMysqlClient() {
	MysqlDemo = mysql.InitMysql(config.Config.MysqlDemo)
}
