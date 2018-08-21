/*基于redis实现的单实例分布式锁*/

package redislock

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/garyburd/redigo/redis"
)

type RedisLock struct {
	lockKey string
	value   string
}

//删除lockkey的lua脚本，保证原子性，
//因为这可以避免误删其他客户端得到的锁，举个例子，
//一个客户端拿到了锁，被某个操作阻塞了很长时间，过了超时时间后自动释放了这个锁，
//然后这个客户端之后又尝试删除这个其实已经被其他客户端拿到的锁。
var delScript = redis.NewScript(1, `if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("del", KEYS[1])
else
	return 0
end`)

//获取锁
func (this *RedisLock) Lock(rd *redis.Conn, timeout int) error {
	{
		//	生成随机数
		b := make([]byte, 16)
		_, err := rand.Read(b)
		if err != nil {
			return err
		}
		this.value = base64.StdEncoding.EncodeToString(b)
	}
	lockReply, err := (*rd).Do("SET", this.lockKey, this.value, "ex", timeout, "nx")
	if err != nil {
		return errors.New("redis exectues fail")
	}
	if lockReply == "OK" {
		return nil
	} else {
		return errors.New("lock fail")
	}
}

//释放锁
func (this *RedisLock) Unlock(rd *redis.Conn) {
	delScript.Do(*rd, this.lockKey, this.value)
}
