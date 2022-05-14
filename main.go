package main

import (
	"elearnping-go/moodle"
	moodlefn "elearnping-go/moodle/function"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Authorization"}
	config.AllowMethods = []string{"GET", "OPTIONS"}
	r.Use(cors.New(config))

	r.GET("/login", func(c *gin.Context) {
		token, invalid := getToken(c)
		if invalid {
			return
		}
		fn := moodlefn.NewFunction[bool](token, moodlefn.VerifyTokenFunction{})
		valid, err := fn.Call()
		if err != nil {
			c.AbortWithStatus(500)
		}
		c.JSON(200, valid)
	})

	r.GET("/sites", func(c *gin.Context) {
		token, invalid := getToken(c)
		if invalid {
			return
		}
		fn := moodlefn.NewFunction[map[string][]moodle.Site](token, moodlefn.GetSitesFunction{
			Classification: "inprogress",
		})
		sites, err := fn.Call()
		if err != nil {
			c.AbortWithStatus(500)
		}
		c.JSON(200, sites)
	})

	r.GET("/sites/:siteid/updates", func(c *gin.Context) {
		token, invalid := getToken(c)
		if invalid {
			return
		}
		siteId, err := strconv.Atoi(c.Param("siteid"))
		if err != nil {
			c.AbortWithStatus(400)
		}
		var from time.Time
		if since, err := strconv.Atoi(c.Query("since")); err != nil {
			from = time.Now().Add(-time.Hour)
		} else {
			from = time.UnixMilli(int64(since))
		}
		fn := moodlefn.NewFunction[moodle.SiteUpdate](token,
			moodlefn.GetUpdatesFunction{
				SiteId: moodle.SiteId(siteId),
				Since:  from,
			},
		)
		updates, err := fn.Call()
		if err != nil {
			c.AbortWithStatus(500)
		}
		c.JSON(200, updates)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

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
