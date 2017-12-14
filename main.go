package main

import (
	"io"
	"os"
)

func main() {
	t := []byte(`package main

import (
	"io"
	"os"
)

func main() {
	t := []byte(X)
	p(t, os.Stdout, true)
}

func p(t []byte, w io.Writer, o bool) {
	for _, b := range t {
		if o && b == 88 {
			w.Write([]byte{96})
			p(t, w, false)
			w.Write([]byte{96})
		} else {
			w.Write([]byte{b})
		}
	}
}
`)
	p(t, os.Stdout, true)
}

func p(t []byte, w io.Writer, o bool) {
	for _, b := range t {
		if o && b == 88 {
			w.Write([]byte{96})
			p(t, w, false)
			w.Write([]byte{96})
		} else {
			w.Write([]byte{b})
		}
	}
}
