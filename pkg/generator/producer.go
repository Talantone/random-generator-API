package generator

import (
	"math/rand"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type Producer struct {
}

func (p *Producer) GenerateRandom(c chan string) {
	if v := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2); v == 1 {
		c <- String(11)
	} else {
		c <- strconv.Itoa(seededRand.Intn(100000000))
	}
}

func NewProducer() *Producer {
	return &Producer{}
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		r := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(charset))
		b[i] = charset[r]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
