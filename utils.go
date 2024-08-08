package idxgorm

import (
	"database/sql"
	"strings"

	"github.com/lunny/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDB 根据提供的数据库驱动和连接字符串初始化并返回一个gorm.DB实例。
// 如果未指定驱动程序，将默认使用MySQL。
// 参数：
//
//	driver - 数据库驱动程序的名称，支持"mysql"和"pgx"。
//	connection - 数据库的连接字符串。
//
// 返回值：
//
//	*gorm.DB - Gorm数据库实例的指针。
func GetDB(driver string, connection string) *gorm.DB {
	// 如果未指定数据库驱动程序，将默认设置为"mysql"。
	if strings.EqualFold(driver, "") {
		driver = "mysql"
	}

	// 初始化错误变量和数据库实例变量。
	var err error
	var x *gorm.DB

	// 根据驱动程序和连接字符串打开SQL数据库连接。
	sqlDB, err := sql.Open(driver, connection)

	// 根据不同的数据库驱动程序初始化gorm数据库实例。
	switch driver {
	case "mysql":
		// 使用MySQL驱动程序初始化gorm数据库实例。
		x, err = gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}), &gorm.Config{})

	case "pgx":
		// 使用pgx驱动程序初始化gorm数据库实例。
		x, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
	}

	// 如果存在错误，记录初始化过程中的错误。
	if nil != err {
		log.Error("init" + err.Error())
	}

	// 返回gorm数据库实例。
	return x
}
