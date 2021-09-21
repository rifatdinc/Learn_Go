package workspace

import (
	"context"
	"log"
	"time"

	"github.com/Ullaakut/nmap/v2"
)

type Returndata struct {
	ports string
	ip    string
}

func Nmap(ip, ports string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(ip),
		nmap.WithPorts(ports),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}
	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}
	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}
	return forHost(result)

}

func forHost(s *nmap.Run) string {
	State := ""
	for _, host := range s.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}
		for _, port := range host.Ports {
			State = port.State.String()
		}
	}
	return State
}
