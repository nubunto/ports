# Ports

Ports is a library designed to build programs that interface directly to a Go program.

## Example

```
package main

import (
	"log"

	"github.com/nubunto/ports"
)

var p = `echo "hello, world!`

func main() {
	program, err := ports.Launch(p, nil)
	if err != nil {
		log.Fatal(err)
	}

	// program is a channel that streams the output from the program,
	// running in a different goroutine
	for out := range program {
		log.Println(string(out))
	}
}

Result: hello, world!
```

This is a work in progress, expect API to change
