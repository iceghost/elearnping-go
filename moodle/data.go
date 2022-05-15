package moodle

import "time"

type ModuleId uint
type SiteId uint
type GroupId int
type Category string
type Classification string

const (
	PAST       Classification = "past"
	INPROGRESS Classification = "inprogress"
	FUTURE     Classification = "future"
)

// a site on moodle, also called "course"
type Site struct {
	Id      SiteId  `json:"id"`
	Name    string  `json:"fullname"`
	GroupId GroupId `json:"groupid"`
}

// module of a site
type Module struct {
	Id   ModuleId `json:"id"`
	Name string   `json:"name"`
	// plugin name, e.g mod/assign, mod/url, mod/quiz, ...
	PluginName string `json:"modname"`
}

// updates of a site in since a timestamp in the past
type SiteUpdate struct {
	Site    Site           `json:"site"`
	From    time.Time      `json:"from"`
	To      time.Time      `json:"to"`
	Updates []ModuleUpdate `json:"updates"`
}

type ModuleUpdate struct {
	Module Module       `json:"module"`
	Areas  []UpdateArea `json:"areas"`
}

// update of a module, i.e what changed
// "configuration" is the most common
//
// TODO: this need improvements
type UpdateArea struct {
	Name string `json:"name"`
}
