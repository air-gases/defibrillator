package defibrillator

import (
	"fmt"

	"github.com/sheng/air"
)

// Gas is used to recover from panics.
func Gas(next air.Handler) air.Handler {
	return func(req *air.Request, res *air.Response) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("%v", r)
			}
		}()

		return next(req, res)
	}
}
