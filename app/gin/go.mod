module github.com/kantaroso/kanta-workspace

go 1.13

require (
	github.com/gin-gonic/gin v1.6.0
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/kantaroso/kanta-workspace/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/kantaroso/kanta-workspace/controllers v0.0.0-00010101000000-000000000000
	github.com/kantaroso/kanta-workspace/lib/accesslog v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1 // indirect
)

replace (
	github.com/kantaroso/kanta-workspace/config => ./config
	github.com/kantaroso/kanta-workspace/controllers => ./controllers
	github.com/kantaroso/kanta-workspace/lib/accesslog => ./lib/accesslog
)
