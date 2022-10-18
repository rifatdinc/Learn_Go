package Datatype

import "fmt"

// burada bir yapi olusturuyoruz
type Vertes struct {
	X int
	Y int
}

func Structs() {
	fmt.Println(Vertes{1, 2})
}

func StructsField() {
	v := Vertes{1, 2}
	//  noktalar ile degerleerine ulasabiliyorz.
	v.X = 3
	v.Y = 6
	fmt.Println(v.X, v.Y)
}

func PointerStruct() {
	v := Vertes{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

var (
	v1 = Vertes{1, 2}  // has type Vertes
	v2 = Vertes{X: 1}  // Y:0 is implicit
	v3 = Vertes{}      // X:0 and Y:0
	p  = &Vertes{1, 2} // has type *vertes
)

func StructInfo() {
	fmt.Println(v1, v2, v3, p)
}
