package gormc

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

type _gorm struct {
	w logger.Writer
}

func (g *_gorm) ConfigPgsql(m *gconfig.Pgsql) *gorm.Config {
	return g.Config(gconfig.DbPgsql, &m.GeneralDB)
}

func (g *_gorm) ConfigMysql(m *gconfig.Mysql) *gorm.Config {
	return g.Config(gconfig.DbMysql, &m.GeneralDB)
}

func (g *_gorm) SetLogger(w logger.Writer) {
	g.w = w
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

	if g.w == nil {
		g.w = NewWriter(dbType, m.LogType, log.New(os.Stdout, "\r\n", log.LstdFlags))
	}
	_default := logger.New(g.w, logger.Config{
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
