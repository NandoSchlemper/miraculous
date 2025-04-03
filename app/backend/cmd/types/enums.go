package types

type Process int

const (
	Anything Process = iota
	Moon
	Earth
	AllPlanets
	Sun
	MilkWay
	AllWorlds
)

type TaskType int

const (
	Ativo TaskType = iota
	Passivo
	Neutralizante
)
