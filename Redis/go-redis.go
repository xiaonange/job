package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "121.37.247.194:6379",
		Password: "Shijian@123456", // no password set
		DB:       0,                // use default DB
	})
    //redis.FailoverOptions{}
	pong, err := client.Ping().Result()
	//sb :=redis.NewFailoverClient( &redis.FailoverOptions{})
	fmt.Println(pong, err)
	//hmset && hset
	//TestHset(client)

	//list
	//TestList(client)

	//set 无序
	//TestSet(client)

	//zset 有序
	//TestZset(client)

	//订阅
	//go TestSubscribe(client)
	//TestPush(client)

	//key和scan
	err = client.Set("aaa","bbb",0).Err()
	err = client.Set("bbb","bbb",0).Err()
	err = client.Set("bcc","bbb",0).Err()
	err = client.Set("acc","bbb",0).Err()
	fmt.Print(err)
	keys,err :=client.Keys("a*").Result()
	fmt.Print(keys,err)
	result,num,err :=client.Scan(0,"",10).Result()
	fmt.Print(result,num,err)


}

func TestString(client *redis.Client) {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func TestHset(client *redis.Client)  {
	type User struct {
		Name string `json:"name"`
		Guid string `json:"guid"`
	}
	//user_info := User{Name: "测试啊",Guid:"ddddd"}

	is_true,err := client.HSet("user_info","name" ,"测试啊").Result()
	fmt.Println(is_true, err)
	result,err :=client.HGet("user_info","name").Result()
	fmt.Print(result)
	//批量
	hmstring,err :=client.HMSet("student", map[string]interface{}{"name":"zlw","guid":"xxxx"}).Result()
	hmstring,err =client.HMSet("student", map[string]interface{}{"name":"like"}).Result()
	fmt.Println(hmstring, err)
	hmreult,err :=client.HMGet("student","name","guid").Result()
	fmt.Println(hmreult, err)
}

func TestList(client *redis.Client) {
	type User struct {
		Name string `json:"name"`
		Guid string `json:"guid"`
	}
	user_info := User{Name: "测试啊",Guid:"ddddd"}
	_, err := client.RPush("list", "sjsjjs", "xxxxx",user_info).Result()
	fmt.Println(err)
	num, err := client.LLen("list").Result()
	fmt.Println(num, err)
	for i := num; i > 0; i-- {
		listRus, err := client.RPop("list").Result()
		fmt.Println(listRus, err)
	}
}

func TestSet(client *redis.Client) {
	result, err := client.SAdd("set", "aaa", "bbb").Result()
	fmt.Println(result, err)
	is_ture, err := client.SIsMember("set", "ccc").Result()
	fmt.Println(is_ture, err)
	data, err := client.SMembers("set").Result()
	fmt.Println(data, err)
}

func TestZset(client *redis.Client)  {
	_, err := client.ZAdd("score", redis.Z{Score: 10, Member: "2222"}).Result()
	_, err = client.ZAdd("score", redis.Z{Score: 9, Member: "333"}).Result()
	result,err :=client.ZRange("score",0,3).Result()
	fmt.Println(result, err)
	result,err =client.ZRevRange("score",0,3).Result()
	fmt.Println(result, err)
	score,err :=client.ZScore("score","333").Result()
	fmt.Println(score, err)
}

func TestPush(client *redis.Client)  {
	err := client.Publish("chat","aaaaa").Err()
	fmt.Println( err)
}

func TestSubscribe(client *redis.Client){
	//参数1 频道名 字符串类型
	pubsub :=client.Subscribe("chat")
	_, err := pubsub.Receive()
	if err != nil {
		return
	}
	ch := pubsub.Channel()
	/*for msg := range ch {
		fmt.Println( msg.Channel, msg.Payload, "\r\n")
	}*/
	for{
		select{
		case message, ok :=<-ch:
			if !ok {
				return
			}
			fmt.Println(message.Payload)
		default:
			fmt.Print("美食家")
		}
	}
}