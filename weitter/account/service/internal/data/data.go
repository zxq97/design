package data

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewDB, NewMemCache, NewAccountRepo)

func NewDB() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:GUOan1992@tcp(127.0.0.1:3306)/zzlove?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return conn
}

func NewMemCache() *memcache.Client {
	return memcache.New([]string{"127.0.0.1:11211"}...)
}
