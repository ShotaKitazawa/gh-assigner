module github.com/ShotaKitazawa/gh-assigner

go 1.13

require (
	cloud.google.com/go/logging v1.0.0
	github.com/ShotaKitazawa/tabemap-api v0.0.0-20191112145853-73fde0b6bb0d // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jinzhu/gorm v1.9.10
	github.com/jmoiron/sqlx v1.2.0
	github.com/nlopes/slack v0.6.1-0.20191106133607-d06c2a2b3249
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.5.0
	github.com/stretchr/testify v1.3.0
	gonum.org/v1/plot v0.0.0-20191107103940-ca91d9d40d0a
	google.golang.org/api v0.7.0
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v1.1.8-0.20190812104308-42bc974514ff
