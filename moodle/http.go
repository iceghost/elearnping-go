package moodle

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"go.uber.org/ratelimit"
)

var rl = ratelimit.New(3)
var logger = log.New(os.Stdout, "[base fetch]", log.Ldate|log.Ltime)

// rate limited fetch function
func Fetch(url url.URL) (body io.ReadCloser, err error) {
	rl.Take()
	logger.Printf("Fetch %s", url.String())
	if res, err := http.Get(url.String()); err != nil {
		return nil, err
	} else {
		return res.Body, nil
	}
}
