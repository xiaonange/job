
package main

import (
"fmt"
"github.com/go-redis/redis"
"time"
)

// 定义redis链接池
var RedisTest *redis.Client

// 初始化redis链接池
func init() {
	RedisTest = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379", // Redis地址
		Password:     "",               // Redis账号
		DB:           1,                // Redis库
		PoolSize:     10,               // Redis连接池大小
		MaxRetries:   3,                // 最大重试次数
		IdleTimeout:  5 * time.Second,  // 空闲链接超时时间
		MinIdleConns: 5,                // 空闲连接数量
	})
	pong, err := RedisTest.Ping().Result()
	if err == redis.Nil {
		fmt.Println("Redis异常")
	} else if err != nil {
		fmt.Println("失败:", err)
	} else {
		fmt.Println(pong)
	}
}

func main() {
	do(true)
	fmt.Println("操作完成")
}

// 传参标记是否必须获取到锁
func do(must bool) {
	res := getLock()
	if res == true {
		// 获取锁成功，需要在业务处理完成后及时删除锁
		defer RedisTest.Del("key")
		// TODO 此处书写业务逻辑

	} else {
		fmt.Println("获取锁失败")
		if must {
			// 如果必须要获取到锁才能往下进行，需要自我休眠
			time.Sleep(100 * time.Millisecond) // 休眠100 毫秒再去执行
			do(true)
		}
	}
}

// 获取锁
func getLock() bool {
	// 这里一定要添加失效时间，以防删除锁失败
	res, err := RedisTest.SetNX("key", true, 500*time.Millisecond).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return res
}
