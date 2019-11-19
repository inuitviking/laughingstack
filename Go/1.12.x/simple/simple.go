package simple

import (
	"fmt"

	"github.com/inuitviking/laughingstack/Go/1.12.x/laugher"
)

var _ laugher.Laugher = (*SimpleLaugher)(nil)

type SimpleLaugher struct{}

// Chuckle laughs twice
func (l SimpleLaugher) Chuckle() {
	fmt.Println(laughN(2))
}

// Chortle laughs five times
func (l SimpleLaugher) Chortle() {
	fmt.Println(laughN(5))
}

// Guffaw laughs ten times
func (l SimpleLaugher) Guffaw() {
	fmt.Println(laughN(10))
}

func laughN(n int) (ha string) {
	for i := 0; i < n; i++ {
		ha += `ha`
	}
	return
}
