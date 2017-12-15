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
	q := "\x60"
	os.Stdout.WriteString(t[:52] + q + t + q + t[53:])
}
`
	q := "\x60"
	os.Stdout.WriteString(t[:52] + q + t + q + t[53:])
}
