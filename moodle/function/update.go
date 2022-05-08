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
		SiteId:  req.SiteId,
		From:    req.Since,
		To:      time.Now(),
		Updates: []moodle.ModuleUpdate{},
	}
	moduleUpdates := make(chan moodle.ModuleUpdate)
	errs := make(chan error)
	for _, instance := range response.Instances {
		go func(instance Instance) {
			fn := GetModuleFunction{Id: instance.Id}
			module, err := NewCachedGetModuleFunction(
				NewFunction[moodle.Module](token, fn),
			).Call()
			if err != nil {
				errs <- err
			} else {
				moduleUpdates <- moodle.ModuleUpdate{
					Module: module,
					Areas:  instance.Areas,
				}
			}
		}(instance)
	}
	for range response.Instances {
		select {
		case update := <-moduleUpdates:
			siteUpdate.Updates = append(siteUpdate.Updates, update)
		case err := <-errs:
			panic(err)
		}
	}
	return siteUpdate
}
