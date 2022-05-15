package query

import (
	"elearnping-go/moodle"
	"encoding/json"
)

type groupsQuery struct{}

// get group ids of a student's site
func NewGroupsQuery() Query[map[moodle.SiteId]moodle.GroupId] {
	return groupsQuery{}
}

func (req groupsQuery) Encode() (string, map[string]string) {
	return "core_group_get_course_user_groups", map[string]string{}
}

func (req groupsQuery) Decode(decoder *json.Decoder) map[moodle.SiteId]moodle.GroupId {
	type Response struct {
		Groups []struct {
			SiteId moodle.SiteId  `json:"courseid"`
			Id     moodle.GroupId `json:"id"`
		} `json:"groups"`
	}
	var response Response
	decoder.Decode(&response)
	output := make(map[moodle.SiteId]moodle.GroupId)
	for _, group := range response.Groups {
		output[group.SiteId] = group.Id
	}
	return output
}
