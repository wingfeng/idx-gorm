package idxgorm

import (
	"database/sql"
	"strings"

	"github.com/lunny/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB(driver string, connection string) *gorm.DB {
	if strings.EqualFold(driver, "") {
		driver = "mysql"
	}
	var err error
	var x *gorm.DB

	sqlDB, err := sql.Open(driver, connection)

	switch driver {
	case "mysql":

		x, err = gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}), &gorm.Config{})

	case "pgx":
		x, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})

	}

	if nil != err {
		log.Error("init" + err.Error())
	}
	return x
}
