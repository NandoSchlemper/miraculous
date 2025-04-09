package domain

type TaskPayload struct {
	Title       string
	Message     string
	Description string
	Talents     float64
	Conciencius float64
	TaskType    string // Ativo, Passivo ou Neutralizante
	// made something to save images, pdfs, etc
}

type EnygmaPayload struct {
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

type TestsPayload struct {
	Title       string
	Message     string
	Description string
	Talents     float64
	Conciencius float64
	ProcessType string // tipos dos processos (0-8)
	// make something to save images, pdfs, etc
}

type UserPayload struct {
	Username string
	Email    string
	Password string
}

type UserAttributesPayload struct {
	TotalConciencius float64
	TotalTalents     float64
	ActualProcess    string // tipos dos processos (0-8)
	ActualTask       string // ID da Task atual
	Score            int    // qnt Tasks feitas até então
}

type UserHistoryPayload struct {
	TaskId   string // ID da task feita
	TaskType string // Enygma, Test ou Task
	Status   bool   // true or false if Task correct
}
