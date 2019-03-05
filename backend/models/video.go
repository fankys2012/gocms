package models

import "time"

type Vod struct {
	Id         string    `xorm:"not null pk autoincr Varchar(32)" json:"id"`
	Name       string    `xorm:"not null Varchar(200)" json:"name"`
	CreateTime time.Time `xorm:"DateTime" json:"create_time"`
	State      int       `xorm:"Int(4)" json:state`
}
