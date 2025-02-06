package repository

import (
	"fmt"
	"time"

	"github.com/Daaaai0809/tamemie/config"
	"github.com/go-sql-driver/mysql"
	g_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DRIVER = "mysql"
	NET    = "tcp"
	PORT   = "3306"
)

func Conn(config *config.Config, host interface{}) (*gorm.DB, error) {
	var dbHost string
	if host == nil {
		dbHost = config.MYSQL_HOST
	} else {
		dbHost = host.(string)
	}

	c := mysql.Config{
		User:                 config.MYSQL_USER,
		Passwd:               config.MYSQL_PASSWORD,
		Net:                  NET,
		Addr:                 fmt.Sprintf("%s:%s", dbHost, PORT),
		DBName:               config.MYSQL_DATABASE,
		AllowNativePasswords: true,
		Params: map[string]string{
			"charset": "utf8mb4",
		},
	}

	count := uint(0)
	db, err := connectDB(c, count)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectDB(dbConf mysql.Config, count uint) (*gorm.DB, error) {
	db, err := gorm.Open(g_mysql.Open(dbConf.FormatDSN()), &gorm.Config{})
	if err != nil {
		if count < 5 {
			time.Sleep(5 * time.Second)
			fmt.Println("failed to connect to database. retrying...")
			connectDB(dbConf, count+1)
		}
		return nil, err
	}

	return db, nil
}
