package main

// import "fmt"
import (
	"elearnping-go/moodle"
	moodlefn "elearnping-go/moodle/function"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("MY_TOKEN")
	r := gin.Default()

	r.GET("/sites", func(c *gin.Context) {
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
		siteId, err := strconv.Atoi(c.Param("siteid"))
		if err != nil {
			c.AbortWithStatus(400)
		}
		fn := moodlefn.NewFunction[moodle.SiteUpdate](token,
			moodlefn.GetUpdatesFunction{
				SiteId: moodle.SiteId(siteId),
				Since:  time.Now().Add(-time.Hour * 24 * 30),
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
