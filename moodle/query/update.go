package query

import (
	"elearnping-go/moodle"
	"encoding/json"
	"fmt"
	"time"
)

type updatesQuery struct {
	Site  moodle.Site
	Since time.Time
}

// get updates of a site, rounded down to nearest 6 hours (0h, 6h, 12h, 18h)
//
// also note that module names are empty, due to only one trip to server
// if you want information about module name, use elearnping-go/moodle/complexquery
func NewUpdatesQuery(site moodle.Site, since time.Time) CachableQuery[moodle.SiteUpdate] {
	return updatesQuery{
		site,
		since.Local().Truncate(6 * time.Hour),
	}
}

func (req updatesQuery) Encode() (string, map[string]string) {
	return "core_course_get_updates_since",
		map[string]string{
			"courseid": fmt.Sprint(req.Site.Id),
			"since":    fmt.Sprint(req.Since.Unix()),
		}
}

func (req updatesQuery) Decode(decoder *json.Decoder) moodle.SiteUpdate {
	type Instance struct {
		Id    moodle.ModuleId     `json:"id"`
		Areas []moodle.UpdateArea `json:"updates"`
	}
	type Response struct {
		Instances []Instance `json:"instances"`
	}
	var response Response
	decoder.Decode(&response)
	siteUpdate := moodle.SiteUpdate{
		Site:    req.Site,
		From:    req.Since,
		To:      time.Now(),
		Updates: make([]moodle.ModuleUpdate, len(response.Instances)),
	}
	for i, instance := range response.Instances {
		siteUpdate.Updates[i] = moodle.ModuleUpdate{
			Module: moodle.Module{Id: instance.Id},
			Areas:  instance.Areas,
		}
	}
	return siteUpdate
}

func (req updatesQuery) CacheKey() string {
	return fmt.Sprintf("updates:%d:%d:%d", req.Site.Id, req.Site.GroupId, req.Since.Unix())
}
