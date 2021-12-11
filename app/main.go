package main

import (
	"github.com/hmrkm/simple-rights/adapter"
	"github.com/hmrkm/simple-rights/io"
	"github.com/hmrkm/simple-rights/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	mysqlDriver "gorm.io/driver/mysql"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	config := Config{}
	if err := envconfig.Process("", &config); err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysqlDriver.Open(io.CreateDSN(
		config.MysqlUser,
		config.MysqlPassword,
		config.MysqlDatabase,
	)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	mysql := io.NewMysql(db)
	defer mysql.Close()

	ru := usecase.NewRights(mysql)
	ra := adapter.NetRights(ru)

	e := echo.New()
	Router(e, ra)
	e.Logger.Fatal(e.Start(":80"))
}
