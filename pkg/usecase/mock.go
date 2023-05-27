package usecase

import (
	"context"
	"github.com/stretchr/testify/mock"
	"random-generator-API/models"
)

type GeneratorUseCaseMock[T models.RandomItemTypes] struct {
	mock.Mock
	typeOfOutput T
}

func (m *GeneratorUseCaseMock[T]) Generate(ctx context.Context, amount *models.Amount, url string) error {
	args := m.Called(ctx, amount, url)

	return args.Error(0)
}

func (m *GeneratorUseCaseMock[T]) GetLastOutput(ctx context.Context, url string) (*models.RandomItem[T], error) {
	args := m.Called(ctx, url)

	return args.Get(0).(*models.RandomItem[T]), args.Error(1)
}
