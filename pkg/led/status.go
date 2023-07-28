package led

type Status struct {
	IsLit    bool   `json:"isLit"`
	Value    uint64 `json:"value"`
	Treshold uint64 `json:"treshold"`
}
