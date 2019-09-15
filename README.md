# Defibrillator

[![GoDoc](https://godoc.org/github.com/air-gases/defibrillator?status.svg)](https://godoc.org/github.com/air-gases/defibrillator)

A useful gas that used to recover from panics for the web applications built
using [Air](https://github.com/aofei/air).

## Installation

Open your terminal and execute

```bash
$ go get github.com/air-gases/defibrillator
```

done.

> The only requirement is the [Go](https://golang.org), at least v1.12.

## Usage

Create a file named `main.go`

```go
package main

import (
	"github.com/air-gases/defibrillator"
	"github.com/aofei/air"
)

func main() {
	a := air.Default
	a.DebugMode = true
	a.Pregases = []air.Gas{
		defibrillator.Gas(defibrillator.GasConfig{}),
	}
	a.GET("/", func(req *air.Request, res *air.Response) error {
		panic("SOME TERRIBLE THINGS HAPPENED!")
		return res.WriteString("This method will never be executed.")
	})
	a.Serve()
}
```

and run it

```bash
$ go run main.go
```

then visit `http://localhost:8080`.

## Community

If you want to discuss Defibrillator, or ask questions about it, simply post
questions or ideas [here](https://github.com/air-gases/defibrillator/issues).

## Contributing

If you want to help build Defibrillator, simply follow
[this](https://github.com/air-gases/defibrillator/wiki/Contributing) to send
pull requests [here](https://github.com/air-gases/defibrillator/pulls).

## License

This project is licensed under the Unlicense.

License can be found [here](LICENSE).
