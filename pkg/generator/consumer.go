package generator

type Consumer struct {
}

func (consumer *Consumer) Consume(m map[int]string, amount int, c chan string) error {

	for i := 1; i < amount+1; i++ {
		m[i] = <-c
	}
	return nil
}

func NewConsumer() *Consumer {
	return &Consumer{}
}
