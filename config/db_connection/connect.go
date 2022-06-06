package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	Driver                  string
	Host                    string
	DbName                  string
	User                    string
	Password                string
	Port                    string
	DBMaxConnections        int
	DBMaxIdleConnection     int
	DBMaxLifetimeConnection int
}

func (c Connection) Conn() (*gorm.DB, error) {
	switch c.Driver {
	case "postgres":
		return c.ConnPG()
	case "mysql":
		return c.ConnMysql()
	default:
		return nil, fmt.Errorf("Driver not found")
	}
}

func (c Connection) ConnPG() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.Host, c.Port, c.User, c.DbName, c.Password)

	gormConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := gormConn.DB()
	if err != nil {
		return nil, err
	}

	db.Ping()
	db.SetMaxOpenConns(c.DBMaxConnections)
	db.SetMaxIdleConns(c.DBMaxIdleConnection)
	db.SetConnMaxLifetime(time.Duration(c.DBMaxLifetimeConnection) * time.Second)

	return gormConn, nil
}

func (c Connection) ConnMysql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.DbName)

	gormConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := gormConn.DB()
	if err != nil {
		return nil, err
	}

	db.Ping()
	db.SetMaxOpenConns(c.DBMaxConnections)
	db.SetMaxIdleConns(c.DBMaxIdleConnection)
	db.SetConnMaxLifetime(time.Duration(c.DBMaxLifetimeConnection) * time.Second)

	return gormConn, nil
}
