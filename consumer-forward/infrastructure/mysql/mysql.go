package mysql

import (
	"database/sql"
	"fmt"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newSQL(cfg Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=%t&loc=%s",
		cfg.User,
		cfg.Password,
		cfg.Endpoint,
		cfg.Name,
		true,
		"Local",
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	size := runtime.NumCPU()
	db.SetMaxIdleConns(size)
	db.SetMaxOpenConns(size*2 + 2)
	db.SetConnMaxLifetime(time.Hour)

	return db
}

func NewORM(sqldb *sql.DB) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		Conn:       sqldb,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		DisableAutomaticPing:   true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return db
}

func CloseSQL(db *gorm.DB) error {
	sqldb, err := db.DB()
	if err != nil {
		return err
	}

	defer sqldb.Close()
	return nil
}

func NewDB(cfg Config) *gorm.DB {
	return NewORM(newSQL(cfg))
}
