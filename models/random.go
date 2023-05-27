package models

type RandomItemTypes interface {
	int | string
}

type RandomItem[T RandomItemTypes] struct {
	Result T `json:"result"`
}
