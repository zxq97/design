package data

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewDB, NewRedis, NewKafkaProducer, NewLikeRepo)

func NewDB() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:GUOan1992@tcp(127.0.0.1:3306)/zzlove_like?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
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

func NewKafkaProducer() sarama.SyncProducer {
	kfkConf := sarama.NewConfig()
	kfkConf.Producer.RequiredAcks = sarama.WaitForAll
	kfkConf.Producer.Retry.Max = 3
	kfkConf.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, kfkConf)
	if err != nil {
		panic(err)
	}
	return producer
}
