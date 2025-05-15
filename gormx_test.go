package gormx

import (
	"fmt"
	"github.com/gomsr/atom-gorm/gconfig"
	"testing"
)

func TestInitDB(t *testing.T) {
	db := InitDB(gconfig.DbServer{
		DbType: gconfig.DbMysql,
		Mysql:  gconfig.Mysql{ /*init*/ },
	})
	fmt.Printf("mysql: %v", db)
}
