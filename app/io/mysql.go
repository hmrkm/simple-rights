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
