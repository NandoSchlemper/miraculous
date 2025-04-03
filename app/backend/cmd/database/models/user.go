package models

import (
	"miraculous/cmd/types"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"unique"`
	Password string
}

type UserAttributes struct {
	gorm.Model
	Conscience    float64
	Talents       float64
	ActualProcess types.Process
	Score         int
	ActualTask    int // task
}

type UserHistory struct {
	gorm.Model
	UserId        string
	TaskId        string
	DataReceive   time.Time
	DataFinish    time.Time
	DataPrevision time.Time
}
