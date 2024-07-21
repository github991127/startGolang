package main

import (
	"context"
	"fmt"
	"time"
	"github.com/go-redis/redis/v8"
)

//Go1.7加入了一个新的标准库context，它定义了Context类型，专门用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作
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

	//设置获取字符串
	// err = rdb.Set(ctx, "username", "zhangsan", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// userinfo, _ := rdb.Get(ctx, "username").Result()
	// fmt.Println(userinfo)

	//设置获取列表

	rdb.LPush(ctx, "hobby", "吃饭", "睡觉")

	// rdb.RPush(ctx, "hobby", "写代码", "打篮球")

	// hoddy, _ := rdb.LRange(ctx, "hobby", 0, -1).Result()

	// fmt.Println(hoddy)

	//设置过期
	rdb.Expire(ctx, "hobby", time.Second*20).Result()

	//设置集合
	// rdb.SAdd(ctx, "hobby", "吃饭", "睡觉")
	// rdb.SAdd(ctx, "hoddy", "吃饭")
	// hoddy, _ := rdb.SMembers(ctx, "hobby").Result()
	// fmt.Println(hoddy)

	// //设置哈希值
	// err2 := rdb.HMSet(ctx, "userinfo", map[string]interface{}{
	// 	"username": "张三",
	// 	"age":      20,
	// }).Err()
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// userinfo, err := rdb.HGetAll(ctx, "userinfo").Result()
	// fmt.Println(userinfo, err)
	// fmt.Println(userinfo["age"])

	go func() {
		time.Sleep(time.Second * 2)
		//发布
		rdb.Publish(ctx, "server0", "this is a message")
	}()

	pubsub := rdb.Subscribe(ctx, "server0")

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}

}

//订阅
