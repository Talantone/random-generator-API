package generator

import (
	"context"

	"random-generator-API/models"
)

type Generator[T models.RandomItemTypes] interface {
	Generate(ctx context.Context, amount *models.Amount) error
	GetLastOutput(ctx context.Context) (*models.RandomItem, error)
}

type RandomGenerator struct {
	Producer *Producer
	Consumer *Consumer
}

func NewRandomGenerator() *RandomGenerator {
	return &RandomGenerator{
		Producer: NewProducer(),
		Consumer: NewConsumer(),
	}
}
