package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func init() {
	sql.Register("mysql", &mysql.MySQLDriver{})
}
