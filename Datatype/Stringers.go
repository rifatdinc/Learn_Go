package Datatype

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func Ipadress() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	fmt.Println(make(chan int, 5))
}
