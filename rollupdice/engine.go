package rollupdice

import (
	"math/rand"
	"time"
)

func newEngine() *engine {
	return &engine{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type engine struct {
	random         *rand.Rand
	value1, value2 int
}

func (e *engine) GenerateValues() {
	e.value1 = e.random.Intn(6) + 1
	e.value2 = e.random.Intn(6) + 1
}

func (e *engine) GetValues() (int, int) { return e.value1, e.value2 }
