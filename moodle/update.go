package moodle

import (
	"encoding/json"
	"fmt"
	"time"
)

type SiteUpdate struct {
	SiteId  SiteId         `json:"siteId"`
	From    time.Time      `json:"from"`
	To      time.Time      `json:"to"`
	Updates []ModuleUpdate `json:"updates"`
}

type ModuleUpdate struct {
	Module Module       `json:"module"`
	Areas  []UpdateArea `json:"areas"`
}

type UpdateArea struct {
	Name string `json:"name"`
}

type GetUpdates struct {
	SiteId SiteId    `json:"siteId"`
	Since  time.Time `json:"since"`
}

func (req GetUpdates) Function() string {
	return "core_course_get_updates_since"
}

func (req GetUpdates) Arguments() map[string]string {
	return map[string]string{"courseid": fmt.Sprint(req.SiteId), "since": fmt.Sprint(req.Since.Unix())}
}

func (req GetUpdates) Decode(decoder *json.Decoder, moodle MoodleService) SiteUpdate {
	type Response struct {
		Instances []struct {
			Id    ModuleId     `json:"id"`
			Areas []UpdateArea `json:"updates"`
		} `json:"instances"`
	}
	var response Response
	decoder.Decode(&response)
	siteUpdate := SiteUpdate{
		SiteId:  req.SiteId,
		From:    req.Since,
		To:      time.Now(),
		Updates: []ModuleUpdate{},
	}
	for _, instance := range response.Instances {
		module, err := Exec[Module](moodle, GetModule{Id: instance.Id})
		if err != nil {
			panic(err)
		}
		siteUpdate.Updates = append(siteUpdate.Updates, ModuleUpdate{
			*module,
			instance.Areas,
		})
	}
	return siteUpdate
}
