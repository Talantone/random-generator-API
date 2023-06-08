package usecase

import (
	"errors"
	"random-generator-API/models"
	"random-generator-API/pkg/generator"
	"sync"
)

const xthreads = 5

var m = make(map[int]string)

type GeneratorUseCase struct {
	generator *generator.RandomGenerator
}

func NewGeneratorUseCase(generator *generator.RandomGenerator) *GeneratorUseCase {
	return &GeneratorUseCase{
		generator: generator,
	}
}

func (g *GeneratorUseCase) Generate(amount *models.Amount) error {
	m = make(map[int]string)
	c := make(chan string, amount.Amount)

	wg := sync.WaitGroup{}
	wg.Add(amount.Amount)
	guard := make(chan struct{}, 100)
	for i := 0; i < amount.Amount; i++ {
		go func(c chan string) {
			guard <- struct{}{}
			defer func() {
				wg.Done()
				<-guard
			}()

			g.generator.Producer.GenerateRandom(c)
		}(c)

	}
	wg.Wait()

	err := g.generator.Consumer.Consume(m, amount.Amount, c)
	if err != nil {
		return err
	}
	return nil
}

func (g *GeneratorUseCase) GetLastOutput(id int) (*models.RandomItem, error) {
	_, ok := m[id]
	item := models.NewRandomItem(m[id])
	if !ok {
		return item, errors.New("element doesn't exist")
	}

	return item, nil
}
