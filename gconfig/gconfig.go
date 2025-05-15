package gconfig

type DbServer struct {
	// DbType 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	DbType string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`

	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
}

//type DbSys struct {
//	DbType string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
//}
