package main

import (
	"fmt"
	//"os"
	"log"

	"github.com/bucknercd/netinfo/pkg/network"
	"github.com/bucknercd/netinfo/pkg/utils"
)

func main() {
	hostname, _ := utils.GetHostname()
	fmt.Printf("Device Name: %s\n\n", hostname)
	interfaces, err := network.GetInterfaces()
	if err == nil {
		for _, i := range interfaces {
			fmt.Println(i)
		}
	}

	fmt.Println("\nChecking internet connectivity ...\n")
	cInfo, e := network.GetConnectivityInfo()
	fmt.Println(cInfo)
	if e != "" {
		log.Fatal(e)
	}


	//get ALL interfaces that are UP and RUNNING
	//get all IP addrs associated with those interfaces (array of Struct)


}
