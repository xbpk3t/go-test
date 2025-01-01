package concs

import (
	"fmt"
	"github.com/sourcegraph/conc/pool"
)

func sema() {
	p := pool.New().WithMaxGoroutines(3)
	for i := 0; i < 10; i++ {
		i := i
		p.Go(func() {
			fmt.Printf("Task %d\n", i)
		})
	}
	p.Wait()
}
