package moodle

import "time"

type ModuleId uint
type SiteId uint
type GroupId int

type Module struct {
	Id         ModuleId `json:"id"`
	Name       string   `json:"name"`
	PluginName string   `json:"modname"`
}

type Site struct {
	Id      SiteId  `json:"id"`
	Name    string  `json:"fullname"`
	GroupId GroupId `json:"groupid"`
}

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

type UpdateArea struct {
	Name string `json:"name"`
}
