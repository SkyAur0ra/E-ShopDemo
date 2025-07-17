// system funtion config(database/jwt/online payment key)
package settings

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Database struct {
	User     string
	Password string
	Host     string
	Name     string
}

var MySQLSetting = &Database{
	User:     "root",
	Password: "1234",
	Host:     "127.0.0.1:3306",
	Name:     "OnlineMall",
}

var Mode = gin.ReleaseMode

var TokenExpireDuration = time.Minute * 30

var Secret = []byte("hello")

var PageSize = 6

var AppId = ""
var AlipayPublicKeyString = ``
var AppPrivateKeyString = ``
