package sessionredis

import (
	"github.com/alexedwards/scs/engine/redisstore"
	"github.com/alexedwards/scs/session"
	"github.com/garyburd/redigo/redis"
	"net/http"
)

func init() {
	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	// Create a new RedisStore instance using the connection pool.
	engine := redisstore.New(pool)

	sessionManager := session.Manage(engine)
}

func Put(r *http.Request, key, val string) bool {
	err := session.PutString(r, key, val)
	if err != nil {
		return false
	}
	return true
}

func Get(r *http.Request, key string) (val string) {
	val, err := session.GetString(r, "message")
	if err != nil {
		val = ""
	}
	return
}
