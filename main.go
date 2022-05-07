package main

// import "fmt"
import (
	"elearnping-go/moodle"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := moodle.BaseMoodleService{Token: os.Getenv("MY_TOKEN")}
	r := gin.Default()

	r.GET("/sites", func(c *gin.Context) {
		sites, err := moodle.Exec[map[string][]moodle.Site](
			&service,
			moodle.GetSites{Classification: "inprogress"},
		)
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
		updates, err := moodle.Exec[moodle.SiteUpdate](
			&service,
			moodle.GetUpdates{
				SiteId: moodle.SiteId(siteId),
				Since:  time.Now().Add(-time.Hour * 24 * 30),
			},
		)
		if err != nil {
			c.AbortWithStatus(500)
		}
		c.JSON(200, updates)
	})
	
	r.Run() // listen and serve on 0.0.0.0:8080
}
