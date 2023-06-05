package usecase

import (
	"log"
	"os"
	"random-generator-API/models"
	"random-generator-API/pkg/generator"
	"sync"
)

type GeneratorUseCase struct {
	generator *generator.RandomGenerator
}

func NewGeneratorUseCase(generator *generator.RandomGenerator) *GeneratorUseCase {
	return &GeneratorUseCase{
		generator: generator,
	}
}

func (g *GeneratorUseCase) Generate(amount *models.Amount, c chan string) error {
	wg := &sync.WaitGroup{}

	for i := 0; i < amount.Amount; i++ {
		wg.Add(1)
		go g.generator.Producer.GenerateRandom(c)
		wg.Done()
	}
	wg.Wait()
	err := g.generator.Consumer.Consume(c)
	if err != nil {
		return err
	}
	return nil
}

func (g *GeneratorUseCase) GetLastOutput() (*models.RandomItem, error) {
	content, err := os.ReadFile("generated.txt")
	if err != nil {
		log.Fatal(err)
	}
	item := models.NewRandomItem(string(content))
	return item, err
}
