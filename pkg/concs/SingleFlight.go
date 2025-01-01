package concs

import (
	"github.com/sourcegraph/conc"
	"sync"
)

var m sync.Map

func SingleFlight(key string, fn func() interface{}) interface{} {
	if val, ok := m.Load(key); ok {
		return val
	}

	var wg conc.WaitGroup
	var res interface{}

	wg.Go(func() {
		res = fn()
		m.Store(key, res)
	})

	wg.Wait()

	return res
}
