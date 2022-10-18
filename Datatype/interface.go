package Datatype

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Write interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Write
}

func Interfaces() {
	var w Write
	fmt.Println(w)
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
