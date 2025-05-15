package internal

import (
	"github.com/gomsr/atom-gorm/gconfig"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Gorm = new(_gorm)

type _gorm struct{}

func (g *_gorm) ConfigPgsql(m *gconfig.Pgsql) *gorm.Config {
	return g.Config(gconfig.DbPgsql, &m.GeneralDB)
}

func (g *_gorm) ConfigMysql(m *gconfig.Mysql) *gorm.Config {
	return g.Config(gconfig.DbMysql, &m.GeneralDB)
}

// Config gorm 自定义配置
func (g *_gorm) Config(dbType string, m *gconfig.GeneralDB) *gorm.Config {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   m.Prefix,
			SingularTable: m.Singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	_default := logger.New(NewWriter(dbType, m.LogType, log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	switch m.LogMode {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
