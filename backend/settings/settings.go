// 系统功能配置（数据库、JWT、在线支付密钥）
package settings

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Database struct {
	User string
	Password string
	Host string
	Name string
}

var MySQLSetting = &Database {
	User: "root",
	Password: "1234",
	Host: "127.0.0.1:3306",
	Name: "OnlineMall",
}

// 环境模式
var Mode = gin.ReleaseMode

// JWT有效时间
var TokenExpireDuration = time.Minute * 30

// JWT加密盐
var Secret = []byte("hello")

// 分页功能，每页数据量
var PageSize = 6

// 支付信息
var AppId = ""
var AlipayPublicKeyString = ``
var AppPrivateKeyString = ``