package generator

import (
	"bufio"
	"os"
)

type Consumer struct {
}

func (consumer *Consumer) Consume(c chan string) error {

	filePath := "generated.txt"
	err := os.Truncate(filePath, 0)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	for ch := 0; ch < len(c); ch++ {
		_, err = w.WriteString(<-c)
		if err != nil {
			return err
		}
		err = w.Flush()
		if err != nil {
			return err
		}
		ch--
	}

	defer func(f *os.File) {
		close(c)
		err := f.Close()
		if err != nil {

		}
	}(f)
	return nil
}

func NewConsumer() *Consumer {
	return &Consumer{}
}
