package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func (d *DatabaseConfig) OpenDatabaseConnection() (*gorm.DB, error) {
	dsn := d.User + ":" + d.Password + "@tcp(" + d.Host + ":" + d.Port + ")/" + d.Database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
