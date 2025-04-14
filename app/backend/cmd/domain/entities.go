package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string
	Message     string
	Description string
	Talents     float64
	Conciencius float64
	TaskType    string // Ativo, Passivo ou Neutralizante
	// made something to save images, pdfs, etc
}

type Enygma struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string
	Message     string
	Description string
	Talents     float64
	Conciencius float64
	Tip1        string
	Tip2        string
	Question    string
	// make something to save images, pdfs, etc
}

type EnygmaAnswers struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	EnygmaId  uuid.UUID `gorm:"foreignKey:EnygmaId;references:ID"`
	Enygma    Enygma    `gorm:"foreignKey:ID"`
	Answer    string
	IsCorrect bool
}

type Tests struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string
	Message     string
	Description string
	Talents     float64
	Conciencius float64
	ProcessType string // tipos dos processos (0-8)
	// make something to save images, pdfs, etc
}

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username string
	Email    string
	Password string
}

type UserAttributes struct {
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TaskId           uuid.UUID `gorm:"foreignKey:TaskId;references:ID"` // ID da Task atual
	Task             Task      `gorm:"foreignKey:TaskId"`
	TaskType         string
	TotalConciencius float64
	TotalTalents     float64
	TotalTaskDone    int
	TotalEnygmaDone  int
	TotalTestDone    int
}

type UserHistory struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserId       uuid.UUID `gorm:"foreignKey:UserId;references:ID"` // id user hur dur
	User         User      `gorm:"foreignKey:UserId"`
	TaskId       uuid.UUID `gorm:"foreignKey:TaskId;references:ID"` // id task hur dur
	Task         uuid.UUID `gorm:"foreignKey:TaskId"`
	Status       bool
	FinishedData time.Time
	StartedData  time.Time
	LimitData    time.Time
	// adicionar alguma coisa para o POST do usuário
}
