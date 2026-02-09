package bootstrap

import (
	"ilicense-lite/bootstrap/client"
	"ilicense-lite/bootstrap/logger"
	"ilicense-lite/bootstrap/otel"
)

func Init(file string) {
	InitConfig(file)
	otel.InitOTEL()
	client.InitMysqlClient()
	logger.InitLogger()
}
