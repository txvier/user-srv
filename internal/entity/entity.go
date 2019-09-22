package entity

import "time"

type Entity struct {
	ID        int64
	CreateAt  time.Time `xorm:"created"`
	UpdateAt  time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}
