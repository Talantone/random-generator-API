package models

type RandomItemTypes interface {
	any
}

type RandomItem struct {
	Result string `json:"result"`
}

func NewRandomItem(result string) *RandomItem {
	return &RandomItem{Result: result}
}
