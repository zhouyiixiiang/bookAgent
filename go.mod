module book_agent

go 1.14

replace (
	basic => ./basic
	config => ./config
	kafka => ./kafka
	model => ./model
	service => ./service
	tcpCs => ./tcpCs
)

require (
	config v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
	kafka v0.0.0-00010101000000-000000000000
	model v0.0.0-00010101000000-000000000000
	service v0.0.0-00010101000000-000000000000
	tcpCs v0.0.0-00010101000000-000000000000
)
