package main

import (
	"elearnping-go/moodle"
	"elearnping-go/moodle/complexquery"
	"elearnping-go/moodle/query"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
	serve()
}

func serve() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Authorization"}
	config.AllowMethods = []string{"GET", "OPTIONS"}
	r.Use(cors.New(config))

	r.GET("/api/login", func(c *gin.Context) {
		token, invalid := getToken(c)
		if invalid {
			return
		}
		valid, err := query.Call(query.NewVerifyQuery(), token)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.JSON(200, valid)
	})

	r.GET("/api/sites", func(c *gin.Context) {
		token, invalid := getToken(c)
		if invalid {
			return
		}
		sites, err := complexquery.CallFullSites(token)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.JSON(200, sites)
	})

	r.GET("/api/updates", func(c *gin.Context) {
		token, invalid := getToken(c)
		if invalid {
			return
		}
		catsites, err := complexquery.CallFullSites(token)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		var from time.Time
		if since, err := strconv.Atoi(c.Query("since")); err != nil {
			from = time.Now()
		} else {
			from = time.UnixMilli(int64(since))
		}
		allUpdates := []moodle.SiteUpdate{}
		for cat := range catsites {
			for _, site := range catsites[cat] {
				updates, err := complexquery.CallFullUpdates(token, site, from)
				if err != nil {
					c.AbortWithError(500, err)
					return
				}
				if len(updates.Updates) > 0 {
					allUpdates = append(allUpdates, updates)
				}
			}
		}
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.JSON(200, allUpdates)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

// extract token from authorization header
func getToken(c *gin.Context) (string, bool) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.String(http.StatusForbidden, "No Authorization header provided")
		c.Abort()
		return "", true
	}
	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		c.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
		c.Abort()
		return "", true
	}
	return token, false
}
