package cache

import (
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_ADDR"),     // use default Addr
	Password: os.Getenv("REDIS_PASSWORD"), // no password set
	DB:       0,                           // use default DB
})

var lock = &sync.Mutex{}
