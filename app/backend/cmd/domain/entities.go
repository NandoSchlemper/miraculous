package domain

type Task struct {
	Title       string
	Message     string
	Description string
	Talents     float64
	Conciencius float64
	TaskType    string // Ativo, Passivo ou Neutralizante
	// made something to save images, pdfs, etc
}

type Enygma struct {
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

type Tests struct {
	Title       string
	Message     string
	Description string
	Talents     float64
	Conciencius float64
	ProcessType string // tipos dos processos (0-8)
	// make something to save images, pdfs, etc
}

type User struct {
	Username string
	Email    string
	Password string
}

type UserAttributes struct {
	TotalConciencius float64
	TotalTalents     float64
	ActualProcess    string // tipos dos processos (0-8)
	ActualTask       string // ID da Task atual
	Score            int
}

type UserHistory struct {
	TaskId   string // ID da task feita
	TaskType string // Enygma, Test ou Task
	Status   bool   // true or false if Task correct
}
