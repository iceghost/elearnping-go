package main

// import "fmt"
import (
	"elearnping-go/moodle"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := moodle.BaseMoodleService{Token: os.Getenv("MY_TOKEN")}
	r := gin.Default()
	r.GET("/sites/:siteid/updates", func(c *gin.Context) {
		siteId, err := strconv.Atoi(c.Param("siteid"))
		if err != nil {
			c.AbortWithStatus(400)
		}
		sites, err := moodle.Exec[moodle.SiteUpdate](
			&service,
			moodle.GetUpdates{
				SiteId: moodle.SiteId(siteId),
				Since:  time.Now().Add(-time.Hour * 24 * 30),
			},
		)
		if err != nil {
			c.AbortWithStatus(500)
		}
		c.JSON(200, sites)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
