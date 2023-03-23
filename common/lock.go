package common

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

type Interface interface {
	Get() bool
	Block(seconds int64) bool
	Release() bool
	ForceRelease()
}

type lock struct {
	context context.Context
	name    string
	owner   string
	seconds int64
}

const releaseLockLuaScript = `
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end
`

func Lock(name string, seconds int64) Interface {
	return &lock{
		context.Background(),
		name,
		uuid.New().String(),
		seconds,
	}
}

func (l *lock) Get() bool {
	return Cache.SetNX(l.context, l.name, l.owner, time.Duration(l.seconds)*time.Second).Val()
}

func (l *lock) Block(seconds int64) bool {
	timer := time.After(time.Duration(seconds) * time.Second)
	for {
		select {
		case <-timer:
			return false
		default:
			if l.Get() {
				return true
			}
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func (l *lock) Release() bool {
	luaScript := redis.NewScript(releaseLockLuaScript)
	result := luaScript.Run(l.context, Cache, []string{l.name}, l.owner).Val().(int64)
	return result != 0
}

func (l *lock) ForceRelease() {
	Cache.Del(l.context, l.name).Val()
}
