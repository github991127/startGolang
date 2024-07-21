package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis数据库连接失败")
	}

	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * 2)
		rdb.Publish(ctx, "server0", "this is a message"+strconv.Itoa(i))
	}

}

//订阅
