package handler

import (
	"random-generator-API/models"
	"random-generator-API/pkg"
)

type Handler[T models.RandomItemTypes] struct {
	useCase pkg.UseCase[T]
}

func NewHandler[T models.RandomItemTypes](useCase pkg.UseCase[T]) *Handler[T] {
	return &Handler[T]{
		useCase: useCase,
	}
}
