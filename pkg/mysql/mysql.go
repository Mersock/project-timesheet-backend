package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	_defaultMaxLifeTime = time.Minute * 3
	_defaultMaxOpenConn = 10
	_defaultMaxIdleConn = 10
)

// NewMysqlConn -.
func NewMysqlConn(url string) (*sql.DB, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, fmt.Errorf("mysql open connection error: %w", err)
	}

	db.SetConnMaxLifetime(_defaultMaxLifeTime)
	db.SetMaxOpenConns(_defaultMaxOpenConn)
	db.SetMaxIdleConns(_defaultMaxIdleConn)

	return db, nil
}
