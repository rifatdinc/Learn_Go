package main

import (
	"Workspace/Datatype"
	s "Workspace/Route_s"
)

func main() {
	// s.Loops("10.50.1.9", "")
	s.Nmap("192.168.14.1", "8728")
	Datatype.Pointer("2")
	Datatype.Structs()
	Datatype.StructsField()
	Datatype.PointerStruct()
	Datatype.StructInfo()
	Datatype.Arrays()
	Datatype.Interfaces()
	Datatype.Ipadress()
}
