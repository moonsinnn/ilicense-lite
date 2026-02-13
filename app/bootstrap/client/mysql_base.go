package client

import (
	"ilicense-lite/config"
	"ilicense-lite/library/mysql"

	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func InitMysqlClient() {
	MysqlDB = mysql.InitMysql(config.Config.MysqlDB)
}
