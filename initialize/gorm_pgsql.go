package initialize

import (
	"github.com/gomsr/atom-gorm/gconfig"
	"github.com/gomsr/atom-gorm/initialize/gormc"
	"github.com/gomsr/atom-gorm/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GormPgSQL 初始化 Postgresql 数据库
func GormPgSQL(p *gconfig.Pgsql) *gorm.DB {
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	db, err := gorm.Open(postgres.New(pgsqlConfig), gormc.Gorm.ConfigPgsql(p))
	if err != nil {
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(p.MaxIdleConns)
	sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	if p.Migration {
		migration.InitializePgsql(&p.GeneralDB, db)
	}

	return db
}

// GormPgSqlByConfig 初始化 Postgresql 数据库 通过参数
func GormPgSqlByConfig(p *gconfig.Pgsql) *gorm.DB {
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}

	db, err := gorm.Open(postgres.New(pgsqlConfig), gormc.Gorm.ConfigPgsql(p))
	if err != nil {
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(p.MaxIdleConns)
	sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	return db
}
