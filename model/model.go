package model

import "gorm.io/gorm"

const(
	TODOMVC_STATUS_ACTIVE    = 0
	TODOMVC_STATUS_COMPLETED = 1
	TODOMVC_STATUS_ALL = -1
)
type Todomvc struct {
	gorm.Model
	Item   string
	Status int
}

type TodomvcAdd struct {
	Item string `json:"item"`
}

type TodomvcDel struct {
	Id int `json:"id"`
}

type TodomvcUpdate struct {
	Id     int   `json:"id"`
	Item   string `json:"item"`
	Status int   `json:"status"`
}

type TodomvcFindStatus struct {
	Status int    `json:"status"`
}

type TodomvcFindItem struct{
	Item   string `json:"item"`
	Status int    `json:"status"`
}
