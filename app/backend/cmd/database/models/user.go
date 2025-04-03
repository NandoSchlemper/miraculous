package models

import (
	"miraculous/cmd/types"
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string
	Email    string `gorm:"unique"`
	Password string
}

type UserAttributes struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Conscience    float64
	Talents       float64
	ActualProcess types.Process
	Score         int
	ActualTask    int // task
}

type UserHistory struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserId        string
	TaskId        string
	DataReceive   time.Time
	DataFinish    time.Time
	DataPrevision time.Time
}
