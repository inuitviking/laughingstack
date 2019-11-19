package concurrent

import (
	"fmt"
	"sync"

	"github.com/inuitviking/laughingstack/Go/1.12.x/laugher"
)

var _ laugher.Laugher = (*ConcurrentLaugher)(nil)

type ConcurrentLaugher struct{}

// Chuckle laughs twice
func (l ConcurrentLaugher) Chuckle() {
	laughN(2)
}

// Chortle laughs five times
func (l ConcurrentLaugher) Chortle() {
	laughN(5)
}

// Guffaw laughs ten times
func (l ConcurrentLaugher) Guffaw() {
	laughN(10)
}

func laugh(wg *sync.WaitGroup) {
	fmt.Print(`ha`)
	wg.Done()
}

func laughN(n int) {
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go laugh(&wg)
	}

	wg.Wait()
	fmt.Printf("\n")
}
