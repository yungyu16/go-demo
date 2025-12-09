package main

import (
	"fmt"
	"net"

	"github.com/labstack/gommon/log"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(fmt.Errorf("djob: could not get network interfaces. %w", err))
	}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			log.Info("djob: could not get network addresses for interface %s", iface.Name)
			continue
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				fmt.Printf("\tIPNet: %v\n", v)
			case *net.IPAddr:
				fmt.Printf("\tIPAddr: %v\n", v)
			}
		}
	}
}
