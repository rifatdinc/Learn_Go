package main

import (
	"Workspace/Datatype"
	"fmt"
)

type Longtimes struct {
	Ints int
	Sd   struct {
		ints int
	}
}

func main() {
	// s.Loops("10.50.1.9", "")
	// s.Nmap("192.168.14.1", "8728")
	// Datatype.Pointer("2")
	// Datatype.Structs()
	// Datatype.StructsField()
	// Datatype.PointerStruct()
	// Datatype.StructInfo()
	// Datatype.Arrays()
	// Datatype.Interfaces()
	// Datatype.Ipadress()
	Datatype.Longtime()
	Longtimes.example(Longtimes{
		Ints: 7,
		Sd:   struct{ ints int }{
			ints: 3,
		},
	})
}

func (c Longtimes) example() {
	fmt.Println(c.Sd.ints)
}
