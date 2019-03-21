package defibrillator

import (
	"fmt"

	"github.com/aofei/air"
)

// GasConfig is a set of configurations for the `Gas`.
type GasConfig struct {
}

// Gas returns an `air.Gas` that is used to recover from panics based on the gc.
func Gas(gc GasConfig) air.Gas {
	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) (err error) {
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
			}()

			return next(req, res)
		}
	}
}
