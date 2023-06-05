package pkg

import (
	"random-generator-API/models"
)

type UseCase interface {
	Generate(amount *models.Amount, c chan string) error
	GetLastOutput() (*models.RandomItem, error)
}
