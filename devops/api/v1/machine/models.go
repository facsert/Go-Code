package machine

import (
	"time"
)

type Machine struct {
	Id        string    `json:"id" gorm:"primaryKey;autoIncrement"`
	Size      int       `json:"size"`
	Locate    string    `json:"locate"`
	CreatedAt time.Time `json:"created_at"`
	Msg       string    `json:"msg"`
	Delete    bool      `json:"delete"`
}
