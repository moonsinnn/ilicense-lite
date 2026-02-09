package mysql

import (
	"io"
	"log"
	"os"
	"time"

	"ilicense-lite/config"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"

	mysql2 "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysql(conf config.MysqlConfig) *gorm.DB {
	InitMysqlLogger(conf)
	c := mysql.Config{
		DriverName:                conf.DataSource.DBDriver,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		DSNConfig: &mysql2.Config{
			DBName:               conf.DataSource.DBName,
			User:                 conf.DataSource.UserName,
			Passwd:               conf.DataSource.Password,
			Addr:                 conf.DataSource.Address,
			AllowNativePasswords: true,
			Net:                  "tcp",
			ParseTime:            true,
			Loc:                  time.Local,
			InterpolateParams:    true,
			Timeout:              time.Duration(conf.DataSource.ConnTimeout) * time.Millisecond,
			ReadTimeout:          time.Duration(conf.DataSource.ReadTimeout) * time.Millisecond,
			WriteTimeout:         time.Duration(conf.DataSource.WriteTimeout) * time.Millisecond,
		},
	}
	if len(conf.DataSource.Params) > 0 {
		m := map[string]string{}
		for _, p := range conf.DataSource.Params {
			m[p.Key] = p.Value
		}
		c.DSNConfig.Params = m
	}
	gdb, err := gorm.Open(mysql.New(c), &gorm.Config{
		Logger: defaultGormLogger,
	})
	if err != nil {
		log.Fatalf("mysql init err: %v", err)
	}
	// 注册 otel 插件
	if err := gdb.Use(otelgorm.NewPlugin()); err != nil {
		log.Fatalf("gdb use otel plugin error: %v", err)
	}
	db, _ := gdb.DB()
	db.SetConnMaxLifetime(time.Duration(conf.DataSource.ConnMaxLifeTime) * time.Millisecond)
	db.SetMaxIdleConns(conf.DataSource.MaxIdleConn)
	db.SetMaxOpenConns(conf.DataSource.MaxOpenConn)
	return gdb
}

func InitMysqlLogger(conf config.MysqlConfig) {
	if defaultGormLogger == nil {
		if conf.LogFile == "" {
			conf.LogFile = "logs/sql.log"
		}
		if conf.LogSlowThreshold == 0 {
			conf.LogSlowThreshold = 1000
		}

		f, _ := os.Create(conf.LogFile)
		defaultGormLogger = logger.New(log.New(io.MultiWriter(f), "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Duration(conf.LogSlowThreshold) * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  conf.Colorful,
		})
	}
}

var defaultGormLogger logger.Interface
