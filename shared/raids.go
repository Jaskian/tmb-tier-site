package shared

var PHASES = []int{1, 2, 3}

type Instance int

const (
	Naxx           = 20
	EoE            = 22
	Sarth          = 24
	UlduarPatterns = 27
	Ulduar         = 28
	Totc10         = 30
	Totc25         = 32
)

var PhaseMappingInstance = map[int]int{
	Naxx:           1,
	EoE:            1,
	Sarth:          1,
	Ulduar:         2,
	UlduarPatterns: 2,
	Totc10:         3,
	Totc25:         3,
}
