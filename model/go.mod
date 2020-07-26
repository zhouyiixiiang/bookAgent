module model

go 1.14

replace config => ../config

require (
	config v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.15
)
