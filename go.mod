module github.com/txvier/user-srv

go 1.12

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.6
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.7.0
	github.com/txvier/user-srv/api v0.0.0-00010101000000-000000000000
	xorm.io/core v0.7.0
)

replace github.com/txvier/user-srv/api => ./api
