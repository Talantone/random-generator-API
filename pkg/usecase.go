package pkg

import (
	"random-generator-API/models"
)

type UseCase interface {
	Generate(amount *models.Amount) error
	GetLastOutput(id int) (*models.RandomItem, error)
}
