package Datatype

import "fmt"

var log = fmt.Println

// Tip cevirme islemi bu sekilde gerceklesiyor.
/*
var i int = 4
var f float64 =float64(i)
var u uint = uint(f)
*/

// or
func TypeConvert() {

	var i int = 4
	var f float64 = float64(i)
	var u uint = uint(f)
	log(i, f, u)
	log(f)
	log(u)

}
