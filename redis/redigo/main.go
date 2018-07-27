// 使用redigo连接redis
//获取 redigo： go get github.com/garyburd/redigo/redis
package main

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "name", "lbbwyt")
	if err != nil {
		fmt.Println("set name filed")
	}
	//需要将取出的值进一步的转换
	name, err := redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", name)
	}

	//判断某个key是否存在
	isExist, err := redis.Bool(c.Do("EXISTS", "name"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", isExist)
	}

	_, err = c.Do("DEL", "mykey")
	if err != nil {
		fmt.Println("redis delelte failed:", err)
	}

	//将map通过json持久化到redis中
	{
		key := "profile"
		imap := map[string]string{
			"username":    "lbbwyt",
			"phoneNumber": "18875251646",
		}
		value, _ := json.Marshal(imap)
		//若给定的 key 已经存在，则 SETNX 不做任何动作。
		n, err := c.Do("SETNX", key, value)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(n)
		}

		var imapGet map[string]string
		valueGet, err := redis.Bytes(c.Do("GET", key))
		if err != nil {
			fmt.Println(err)
		}
		errShal := json.Unmarshal(valueGet, &imapGet)
		if errShal != nil {
			fmt.Println(err)
		}
		fmt.Println(imapGet["username"])
		fmt.Println(imapGet["phoneNumber"])
	}
	//	设置过期时间
	{
		key := "profile"
		n, err := c.Do("EXPIRE", key, 60)
		if err != nil {
			fmt.Println("set EXPIRE fail")
		} else {
			fmt.Println(n)
		}
	}
	//列表操作
	{
		_, err = c.Do("lpush", "runoobkey", "redis")
		if err != nil {
			fmt.Println("redis set failed:", err)
		}

		_, err = c.Do("lpush", "runoobkey", "mongodb")
		if err != nil {
			fmt.Println("redis set failed:", err)
		}

		values, _ := redis.Values(c.Do("lrange", "runoobkey", "0", "100"))
		for _, v := range values {
			fmt.Println(string(v.([]byte)))
		}
	}

	//管道操作
	//	客户端可以发送多个命令到服务器而无需等待响应，最后在一次读取多个响应
	{
		c.Send("SET", "name", "wyt")
		c.Send("SET", "name1", "lbb")
		c.Send("GET", "name")
		c.Flush()
		c.Receive()
		v, err := redis.String(c.Receive())
		if err != nil {
			fmt.Println("reply from GET failed")
		} else {
			fmt.Println(v)
		}
	}

	//scan函数能将数组转换成go类型
	var value1 string
	var value2 string
	reply, err := redis.Values(c.Do("MGET", "name", "name1"))
	if err != nil {
		fmt.Println("reply from MGET failed")
	}
	if _, err := redis.Scan(reply, &value1, &value2); err != nil {
		fmt.Println("convert reply from MGET failed")
	} else {
		fmt.Println("&&&" + value1 + value2)
	}

	//事务处理
	//使用索引号作为分数//这里使用管道添加测试数据
	{
		for i, member := range []string{"red", "blue", "green"} {
			c.Send("ZADD", "zset", i, member)
		}

		if _, err := c.Do(""); err != nil {
			fmt.Println(err)
			return
		}
		//使用事务从zset中删除元素
		v, err := zpop(c, "zset")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v)
		//使用脚本
		v, err = redis.String(zpopScript.Do(c, "zset"))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v)

		//使用脚本（1）
		v, err = redis.String(zpopScript.Do(c, "zset"))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v)
		//使用脚本（2）
		v, err = redis.String(zpopScript.Do(c, "zset"))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v)

	}

}

// zpop pops a value from the ZSET key using WATCH/MULTI/EXEC commands.
var zpopScript = redis.NewScript(1, `
	local r = redis.call('ZRANGE', KEYS[1],0,0)
	if r ~= nil then 
	 	r = r[1]
		redis.call('ZREM', KEYS[1], r)
	end 
	return r
`)

// zpop pops a value from the ZSET key using WATCH/MULTI/EXEC commands.
func zpop(c redis.Conn, key string) (result string, err error) {
	defer func() {
		if err != nil {
			c.Do("DISCARD")
		}
	}()
	//循环直到事务成功提交
	for {
		if _, err := c.Do("WATCH", key); err != nil {
			return "", err
		}
		members, err := redis.Strings(c.Do("ZRANGE", key, 0, 0))
		if err != nil {
			return "", err
		}
		if len(members) != 1 {
			return "", redis.ErrNil
		}
		c.Send("MULTI")
		c.Send("ZREM", key, members[0])
		queued, err := c.Do("EXEC")
		if err != nil {
			return "", err
		}
		if queued != nil {
			result = members[0]
			break
		}

	}
	return result, nil

}
