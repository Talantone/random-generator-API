package pkg

import (
	"context"
	"random-generator-API/models"
)

type UseCase[T models.RandomItemTypes] interface {
	Generate(ctx context.Context, amount *models.Amount, url string) error
	GetLastOutput(ctx context.Context, url string) (*models.RandomItem[T], error)
}
