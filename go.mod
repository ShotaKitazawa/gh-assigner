module github.com/ShotaKitazawa/gh-assigner

go 1.13

require (
	cloud.google.com/go/logging v1.0.0
	github.com/ShotaKitazawa/tabemap-api v0.0.0-20191112145853-73fde0b6bb0d
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/jinzhu/gorm v1.9.10
	github.com/jmoiron/sqlx v1.2.0
	github.com/nlopes/slack v0.6.0
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.5.0
	github.com/stretchr/testify v1.3.0
	gopkg.in/go-playground/assert.v1 v1.2.1
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
