package components

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var RDb *sql.DB
var WDb *sql.DB

const MaxIdle int = 20
const MaxOpen int = 20

// dsn = "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4"

func InitMysql(rDsn, wDsn string) error {
	if db, err := sql.Open("mysql", rDsn); err == nil {
		db.SetConnMaxLifetime(time.Second * 20)
		db.SetMaxIdleConns(MaxIdle)
		db.SetMaxOpenConns(MaxOpen)
		if err := db.Ping(); err == nil {
			RDb = db
		} else {
			return err
		}
	} else {
		return err
	}
	if db, err := sql.Open("mysql", wDsn); err == nil {
		db.SetConnMaxLifetime(time.Second * 20)
		db.SetMaxIdleConns(MaxIdle)
		db.SetMaxOpenConns(MaxOpen)
		if err := db.Ping(); err == nil {
			WDb = db
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}
