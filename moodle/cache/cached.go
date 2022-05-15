package cache

import (
	"bytes"
	"context"
	"elearnping-go/moodle"
	"elearnping-go/moodle/query"
	"encoding/json"
	"io"
	"strings"
	"time"
)

var ctx = context.Background()

// call queries and cache them
func Call[Res any](q query.CachableQuery[Res], token string) (Res, error) {
	defer lock.Unlock()
	lock.Lock()

	var body io.ReadCloser

	if v, err := rdb.Get(ctx, q.CacheKey()).Result(); err == nil {
		// get from cache first
		body = io.NopCloser(strings.NewReader(v))
	} else if b, err := moodle.Fetch(query.URL[Res](q, token)); err == nil {
		// if cached body not found, fetch fresh body and copy into cache

		var buf bytes.Buffer
		defer func() {
			rdb.SetNX(ctx, q.CacheKey(), buf.String(), 6*time.Hour)
		}()

		// when body is read, data is copied into buffer for defered caching
		body = io.NopCloser(io.TeeReader(b, &buf))
	} else {
		var zero Res
		return zero, err
	}

	defer body.Close()
	decoder := json.NewDecoder(body)
	return q.Decode(decoder), nil
}
