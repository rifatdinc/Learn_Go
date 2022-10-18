package workspace

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"gopkg.in/routeros.v2"
)

var (
	async  = flag.Bool("async", false, "Use async code")
	useTLS = flag.Bool("tls", false, "Use TLS")
)
var wg sync.WaitGroup

func dial(address, username, password string) (*routeros.Client, error) {
	if *useTLS {
		return routeros.DialTLS(address, username, password, nil)
	}
	return routeros.Dial(address, username, password)
}

func command(address, username, password string) routeros.Reply {
	flag.Parse()

	c, err := dial(address, username, password)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	if *async {
		c.Async()
	}

	r, err := c.Run("/system/identity/print")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(r)
	return *r
}
func Loops(user string, passwd string, ip ...string) {
	defer wg.Done()

	for _, v := range ip {

		fmt.Println(command(""+v+":8728", user, passwd))
	}
}
