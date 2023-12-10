package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	v1 "github.com/zxq97/design/fuxi/api/service/v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewDB, NewRedis, NewFuxiBFFRepo, NewFuxiClient)

func NewDB() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:GUOan1992@tcp(127.0.0.1:3306)/zzlove_fuxi?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return conn
}

func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		DB:           0,
		DialTimeout:  time.Second * 5,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	})
}

func NewFuxiClient() v1.FuxiClient {
	var opts = []grpc.ClientOption{
		grpc.WithEndpoint(":8889"),
		grpc.WithTimeout(time.Millisecond * 200),
		grpc.WithMiddleware(recovery.Recovery()),
	}

	conn, err := grpc.DialInsecure(context.Background(), opts...)
	if err != nil {
		panic(err)
	}
	return v1.NewFuxiClient(conn)
}
