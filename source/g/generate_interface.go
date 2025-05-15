package g

// https://gorm.io/gen/sql_annotation.html

type upsTagInterface interface {

	// select id from @@table where id=@id
	FindNameById(id int) int64
}
