package io

import (
	"fmt"
)

type Mysql struct {
	conn GormConn
}

func NewMysql(conn GormConn) Mysql {
	return Mysql{
		conn: conn,
	}
}

func CreateDSN(user string, password string, database string) (dsn string) {
	return fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, database)
}

func (m Mysql) Close() error {
	db, err := m.conn.DB()
	if err != nil {
		return err
	}

	db.Close()

	return nil
}

func (m Mysql) Load(destAddr interface{}, cond interface{}) error {
	return m.conn.Where(cond).Find(destAddr).Error
}

func (m Mysql) LoadWith(destAddr interface{}, destCond interface{}, withModelName string, withCond interface{}) error {
	if destCond == nil {
		if withCond == nil {
			return m.conn.Joins(withModelName).Find(destAddr).Error
		}
		return m.conn.Joins(withModelName, m.conn.Where(withCond)).Find(destAddr).Error
	}
	if withCond == nil {
		return m.conn.Joins(withModelName).Where(destCond).Find(destAddr).Error
	}
	return m.conn.Joins(withModelName, m.conn.Where(withCond)).Where(destCond).Find(destAddr).Error
}
