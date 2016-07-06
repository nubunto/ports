# Ports

Ports is a library designed to build programs that interface directly to a Go program.

## Example

```
package main

import (
	"fmt"

	"github.com/nubunto/ports"
)

func main() {
	program, err := ports.Launch("echo", []string{"hello, world!"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(program))
}

Result: hello, world!
```

This is a work in progress, expect API to change
