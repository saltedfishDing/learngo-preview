package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"preview/conf"
)

var (
	DB *gorm.DB
)

func Init(m *conf.Database) error {

	var (
		db  *gorm.DB
		err error
	)

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)

	DB = db
	return nil
}
