package main

import (
	"os"
)

func main() {
	t := `package main

import (
	"os"
)

func main() {
	t := _
	q := "\u0060"
	os.Stdout.WriteString(t[:52] + q + t + q + t[53:])
}
`
	q := "\u0060"
	os.Stdout.WriteString(t[:52] + q + t + q + t[53:])
}
