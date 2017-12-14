package main

import (
	"io"
	"os"
)

type B []byte

func main() {
	t := B(`package main

import (
	"io"
	"os"
)

type B []byte

func main() {
	t := B(X)
	p(t, os.Stdout, 2)
}
func p(t B, w io.Writer, o byte) {
	for _, b := range t {
		if o+b == 90 {
			w.Write(B{96})
			p(t, w, 0)
			w.Write(B{96})
		} else {
			w.Write(B{b})
		}
	}
}
`)
	p(t, os.Stdout, 2)
}
func p(t B, w io.Writer, o byte) {
	for _, b := range t {
		if o+b == 90 {
			w.Write(B{96})
			p(t, w, 0)
			w.Write(B{96})
		} else {
			w.Write(B{b})
		}
	}
}
