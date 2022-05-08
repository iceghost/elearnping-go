package function

import (
	"elearnping-go/moodle"
	"encoding/json"
	"fmt"
	"time"
)

type GetUpdatesFunction struct {
	SiteId moodle.SiteId `json:"siteId"`
	Since  time.Time     `json:"since"`
}

func (req GetUpdatesFunction) Name() string {
	return "core_course_get_updates_since"
}

func (req GetUpdatesFunction) Arguments() map[string]string {
	return map[string]string{"courseid": fmt.Sprint(req.SiteId), "since": fmt.Sprint(req.Since.Unix())}
}

func (req GetUpdatesFunction) Decode(token string, decoder *json.Decoder) moodle.SiteUpdate {
	type Response struct {
		Instances []struct {
			Id    moodle.ModuleId     `json:"id"`
			Areas []moodle.UpdateArea `json:"updates"`
		} `json:"instances"`
	}
	var response Response
	decoder.Decode(&response)
	siteUpdate := moodle.SiteUpdate{
		SiteId:  req.SiteId,
		From:    req.Since,
		To:      time.Now(),
		Updates: []moodle.ModuleUpdate{},
	}
	for _, instance := range response.Instances {
		fn := GetModuleFunction{Id: instance.Id}
		module, err := NewFunction[moodle.Module](token, fn).Call()
		if err != nil {
			panic(err)
		}
		siteUpdate.Updates = append(siteUpdate.Updates, moodle.ModuleUpdate{
			module,
			instance.Areas,
		})
	}
	return siteUpdate
}
