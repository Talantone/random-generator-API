package usecase

import (
	"github.com/stretchr/testify/mock"
	"random-generator-API/models"
)

type GeneratorUseCaseMock struct {
	mock.Mock
}

func (m *GeneratorUseCaseMock) Generate(amount *models.Amount, c chan string) error {
	args := m.Called(amount)

	return args.Error(0)
}

func (m *GeneratorUseCaseMock) GetLastOutput() (*models.RandomItem, error) {
	args := m.Called()

	return args.Get(0).(*models.RandomItem), args.Error(1)
}
