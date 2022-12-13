# greg

[![Build Status](https://github.com/hebcal/greg/actions/workflows/go.yml/badge.svg)](https://github.com/hebcal/greg/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hebcal/greg)](https://goreportcard.com/report/github.com/hebcal/greg)
[![GoDoc](https://pkg.go.dev/badge/github.com/hebcal/greg?status.svg)](https://pkg.go.dev/github.com/hebcal/greg)

Hebcal's greg package converts between Gregorian dates
and R.D. (Rata Die) day numbers.

Example

```golang
package main

import (
	"fmt"
	"time"

	"github.com/hebcal/greg"
)

func main() {
	rataDie := greg.ToRD(1995, time.December, 17)
	fmt.Println(rataDie)
	// Output: 728644
}
```
