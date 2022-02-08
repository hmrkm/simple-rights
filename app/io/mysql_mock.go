package io

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MockTable struct {
	ID             string
	Name           string
	MockSubTableID string
}

func NewMysqlMock() (Mysql, sqlmock.Sqlmock) {
	db, sqlMock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(
		mysql.New(
			mysql.Config{
				Conn: db,
			}), &gorm.Config{})
	mysql := NewMysql(gormDB)
	return mysql, sqlMock
}
