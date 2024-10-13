package main

import (
	"fmt"
	"log"

	"github.com/bucknercd/netinfo/pkg/network"
	"github.com/bucknercd/netinfo/pkg/utils"
)

func main() {
	hostname, _ := utils.GetHostname()
	fmt.Printf("Device Name: %s\n\n", hostname)
	fmt.Println("===== Interface List =====")
	interfaces, err := network.GetInterfaces()
	if err == nil {
		for _, i := range interfaces {
			fmt.Println(i)
		}
	}

	fmt.Println("\nChecking internet connectivity ...\n")
	cInfo, err := network.GetConnectivityInfo()
	fmt.Println(cInfo)
	if err != nil {
		log.Fatal(err)
	}
}
