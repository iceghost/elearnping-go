package function

import (
	"bytes"
	"context"
	"elearnping-go/moodle"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_ADDR"),     // use default Addr
	Password: os.Getenv("REDIS_PASSWORD"), // no password set
	DB:       0,                           // use default DB
})

type CachedGetModuleFunction struct {
	base  BaseFunction[moodle.Module]
	mutex *sync.Mutex
}

func NewCachedGetModuleFunction(base BaseFunction[moodle.Module]) CachedGetModuleFunction {
	return CachedGetModuleFunction{
		base,
		&sync.Mutex{},
	}
}

var ctx = context.Background()

func key[Res any](fn Function[Res]) string {
	return fmt.Sprintf("%s:%v", fn.Name(), fn.Arguments())
}

func (fn CachedGetModuleFunction) Fetch() (io.ReadCloser, error) {
	defer fn.mutex.Unlock()
	fn.mutex.Lock()
	if v, err := rdb.Get(ctx, key(fn.base.iFn)).Result(); err == nil {
		// get from cache first
		return io.NopCloser(strings.NewReader(v)), nil
	} else {
		// else fresh fetch
		return fn.base.Fetch()
	}
}

func (fn CachedGetModuleFunction) Call() (moodle.Module, error) {
	body, err := fn.Fetch()
	if err != nil {
		return moodle.Module{}, err
	}
	var buf bytes.Buffer
	defer func() {
		body.Close()
		rdb.SetNX(ctx, key(fn.base.iFn), buf.String(), time.Hour)
	}()
	decoder := json.NewDecoder(io.TeeReader(body, &buf))
	return fn.base.iFn.Decode(fn.base.token, decoder), nil
}
