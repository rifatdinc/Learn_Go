package Datatype

import "fmt"

func Pointer(s string) {
	//  i ve j nin degeleri
	i, j := 42, 2701
	// i nin ram adresine  p yi atiyoruz
	p := &i
	fmt.Println(*p)
	// burada basilan deger 42
	*p = 21
	// burada ramdaki adresi degistiriyoruz.
	fmt.Println(i)
	// burada basilan deger ise 21

	x := &j
	*x = *x / 37
	// burada j nin degeride 73
	fmt.Println(x)
}