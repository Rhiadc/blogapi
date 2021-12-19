package redis

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	//"github.com/twinj/uuid"
)

var Client *redis.Client

func InitRedis() {
	//Initializing redis

	fmt.Println("Testando")
	dsn := "redis:6379"
	Client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(Client)
}
