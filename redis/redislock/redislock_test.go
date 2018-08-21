// redislock_test
package redislock

import (
	"fmt"
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
)

var Redispool *redis.Pool

//初始化连接池
func init() {
	Redispool = &redis.Pool{
		//最大闲置连接数
		MaxIdle: 10,
		//闲置300秒后关闭连接
		IdleTimeout: 300 * time.Second,
		//创建和配置连接
		Dial: func() (redis.Conn, error) {
			tcp := fmt.Sprintf("%s:%d", "127.0.0.1", 6379)
			c, err := redis.Dial("tcp", tcp)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		//可选的选项，在连接前去检查闲置连接的状态
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func TestLock(t *testing.T) {
	rd := Redispool.Get()
	defer rd.Close()
	go func() {
		Alock := RedisLock{lockKey: "XXX"}
		err := Alock.Lock(&rd, 5)   //5 秒后自动删除Alock
		time.Sleep(7 * time.Second) //等待7秒
		fmt.Println("111", err)
		Alock.Unlock(&rd) //想删除的是Alock锁，但是Alock 已经被自动删除 ,Block由于value 不一样，所以也不会删除
	}()
	time.Sleep(6 * time.Second) //此时Alock 已经被删除
	Block := RedisLock{lockKey: "XXX"}
	err := Block.Lock(&rd, 5) //此时 会获取新的lock Block
	fmt.Println("222", err)

	time.Sleep(2 * time.Second)
	Clock := RedisLock{lockKey: "XXX"}
	err = Clock.Lock(&rd, 5) //想获取新的lock Clock，但由于 Block还存在，返回错误
	fmt.Println("333", err)

	time.Sleep(10 * time.Second)
}
