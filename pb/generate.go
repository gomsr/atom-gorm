package pb

import (
	"database/sql"
	"github.com/Mikaelemmmm/sql2pb/core"
	"github.com/alice52/jasypt-go"
	"github.com/kongmsr/atom-gorm/gconfig"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

const (
	dbType     = "mysql"
	pkg        = "pb"
	goPkg      = "./pb"
	table      = "*"
	fieldStyle = "sql_pb"
)

// Deprecated: G0
func G0(service, dbType, mdsn string) error {
	if v, err := jasypt.New().Decrypt(mdsn); err == nil {
		mdsn = v
	}

	db, err := sql.Open(dbType, mdsn)
	if err != nil {
		return err
	}
	defer db.Close()

	return Gpb(db, service, pkg, goPkg, table, fieldStyle, nil, nil)
}

func G(service string, m *gconfig.DbServer) error {
	switch m.DbType {
	case gconfig.DbMysql:
		return GM(service, &m.Mysql)
	case gconfig.DbPgsql:
		return GP(service, &m.Pgsql)
	default:
		panic("unknown db type")
	}
}

func GP(service string, m *gconfig.Pgsql) error {
	gdb, err := gorm.Open(postgres.Open(m.Dsn()))
	if err != nil {
		return err
	}
	db, err := gdb.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	return Gpb(db, service, pkg, goPkg, table, fieldStyle, nil, nil)
}

func GM(service string, m *gconfig.Mysql) error {

	gdb, err := gorm.Open(mysql.Open(m.Dsn()))
	if err != nil {
		return err
	}
	db, err := gdb.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	return Gpb(db, service, pkg, goPkg, table, fieldStyle, nil, nil)
}

func Gpb(db *sql.DB, service, pkg, goPkg, table, fieldStyle string, ignoreTables, ignoreColumns []string) error {

	s, err := core.GenerateSchema(db, table, ignoreTables, ignoreColumns, service, goPkg, pkg, fieldStyle)
	if nil != err {
		return err
	}

	filePath := "./pb/" + service + ".proto"
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(filePath, []byte(s.String()), 0644); err != nil {
		return err
	}

	return nil
}
