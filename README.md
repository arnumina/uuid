[![GoDoc](https://godoc.org/github.com/arnumina/uuid?status.svg)](https://godoc.org/github.com/arnumina/uuid)
[![Go Report Card](https://goreportcard.com/badge/github.com/arnumina/uuid)](https://goreportcard.com/report/github.com/arnumina/uuid)
![CI](https://github.com/arnumina/uuid/workflows/CI/badge.svg)

# uuid

This module allows you to generate random strings in UUID format.
It is not intended to be RFC compliant, merely to use a well-understood string representation of a 128-bit value.

## Example

```go
package main

import (
	"fmt"

	"github.com/arnumina/uuid"
)

func main() {
	id := uuid.New()
	fmt.Printf("UUID: %s\n", id)
}
```
```
UUID: df6ace8a-a66b-4c2e-5e9b-5d5508e92c2c
```

---
Copyright (c) 2020 Institut National de l'Audiovisuel
