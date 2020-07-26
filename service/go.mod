module service

go 1.14

replace (
	basic => ../basic
	config => ../config
	kafka => ../kafka
	tcpCs => ../tcpCs
)

require (
	config v0.0.0-00010101000000-000000000000
	kafka v0.0.0-00010101000000-000000000000
	tcpCs v0.0.0-00010101000000-000000000000
)
