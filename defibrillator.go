package defibrillator

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/aofei/air"
)

// stackPool is the pool of stack messages.
var stackPool = &sync.Pool{
	New: func() interface{} {
		return make([]byte, 4<<10)
	},
}

// GasConfig is a set of configurations for the `Gas`.
type GasConfig struct {
	DisableIncludeStacks bool

	Skippable func(*air.Request, *air.Response) bool
}

// Gas returns an `air.Gas` that is used to recover from panics based on the gc.
func Gas(gc GasConfig) air.Gas {
	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) (err error) {
			if gc.Skippable != nil && gc.Skippable(req, res) {
				return next(req, res)
			}

			defer func() {
				r := recover()
				if r == nil {
					return
				}

				var isError bool
				err, isError = r.(error)
				if !isError {
					err = fmt.Errorf("%v", r)
				}

				if gc.DisableIncludeStacks {
					return
				}

				stack := stackPool.Get().([]byte)
				length := runtime.Stack(stack, true)
				err = fmt.Errorf("%v: %s", err, stack[:length])
				stackPool.Put(stack)
			}()

			return next(req, res)
		}
	}
}
