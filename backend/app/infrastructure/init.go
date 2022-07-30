package infrastructure

import (
	"fmt"

	"github.com/jumpei00/board/backend/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GenerateDBPool() (*gorm.DB, error) {
	// dsn -> username:password@protocol(host:port)/dbname?param=value
	dsn := fmt.Sprintf(
		"%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True",
		config.GetMySQLUserName(),
		config.GetMySQLPassword(),
		config.GetMysqlProtocol(),
		config.GetMySQLHost(),
		config.GetMySQLDatabaseName(),
	)

	dbLogger := logger.Default
	// 本番環境の場合はDBのログを出力しない
	if config.IsProduction() {
		dbLogger.LogMode(logger.Silent)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dbLogger})
	if err != nil {
		return nil, err
	}

	return db, nil
}