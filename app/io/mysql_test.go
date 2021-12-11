package io

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"gorm.io/gorm"
)

func TestCreateDSN(t *testing.T) {
	testCases := []struct {
		name     string
		user     string
		password string
		database string
		expected string
	}{
		{
			"正常ケース",
			"user",
			"passwd",
			"db",
			"user:passwd@tcp(mysql:3306)/db?charset=utf8mb4&parseTime=True&loc=Local",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := CreateDSN(tc.user, tc.password, tc.database)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("CreateDSN() value is missmatch :%s", diff)
			}
		})
	}
}

func TestClose(t *testing.T) {
	mysql, _ := NewMysqlMock()
	testCases := []struct {
		name     string
		msql     Mysql
		err      error
		expected error
	}{
		{
			"正常ケース",
			mysql,
			nil,
			nil,
		},
		{
			"異常ケース",
			mysql,
			gorm.ErrInvalidDB,
			gorm.ErrInvalidDB,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mysql := tc.msql
			if tc.err != nil {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				mdbc := NewMockGormConn(ctrl)
				mdbc.EXPECT().DB().Return(&sql.DB{}, tc.err)
				mysql.conn = mdbc
			}

			actual := mysql.Close()

			if !errors.Is(actual, tc.expected) {
				t.Errorf("Close() actualErr: %v, ecpectedErr: %v", actual, tc.expected)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	testCases := []struct {
		name        string
		cond        MockTable
		dbID        string
		dbName      string
		expected    MockTable
		expectedErr error
	}{
		{

			"正常ケース",
			MockTable{ID: "1"},
			"1",
			"aaa",
			MockTable{
				ID:   "1",
				Name: "aaa",
			},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mysql, sqlMock := NewMysqlMock()

			actual := MockTable{}
			sqlMock.ExpectQuery(regexp.QuoteMeta(
				"SELECT * FROM `mock_tables` WHERE `mock_tables`.`id` = ?",
			)).
				WithArgs(tc.cond.ID).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
					AddRow(tc.dbID, tc.dbName))

			actualErr := mysql.Load(&actual, tc.cond)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("Find() value is missmatch :%s", diff)
			}
			if !errors.Is(actualErr, tc.expectedErr) {
				t.Errorf("Find() actualErr: %v, ecpectedErr: %v", actualErr, tc.expectedErr)
			}
		})
	}
}
