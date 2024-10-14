package network

import (
	"fmt"
	"net"
	"strings"
)

type CustomInterface struct {
	Name           string
	HwAddr         string
	Status         string
	IPAddr         string
	Netmask        string
	DefaultGateway string
}

func (i *CustomInterface) String() string {
	var out string
	out += fmt.Sprintf("%-20s %-10s\n", "Interface Name:", i.Name)
	if i.HwAddr != "" {
		out += fmt.Sprintf("%-20s %-10s\n", "Hardware Address:", i.HwAddr)
	}

	out += fmt.Sprintf("%-20s %-10s\n", "Interface Status:", i.Status)

	if i.IPAddr != "" {
		out += fmt.Sprintf("%-20s %-10s\n", "IP Address:", i.IPAddr)
	}

	if i.Netmask != "" {
		out += fmt.Sprintf("%-20s %-10s\n", "Subnet Mask:", i.Netmask)
	}
	return out
}

func toIPv4(mask string) string {
	maskMap := map[string]string{"16": "255.255.0.0",
		"17": "255.255.128.0",
		"18": "255.255.192.0",
		"19": "255.255.224.0",
		"20": "255.255.240.0",
		"21": "255.255.248.0",
		"22": "255.255.252.0",
		"23": "255.255.254.0",
		"24": "255.255.255.0",
		"25": "255.255.255.128",
		"26": "255.255.255.192",
		"27": "255.255.255.224",
		"28": "255.255.255.240",
		"29": "255.255.255.248",
		"30": "255.255.255.252",
		"31": "255.255.255.254"}
	return maskMap[mask]
}

func GetInterfaces() ([]*CustomInterface, error) {
	var interfs []*CustomInterface
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range interfaces {
		var status []string
		inter := &CustomInterface{}
		inter.Name = i.Name
		inter.HwAddr = i.HardwareAddr.String()

		if i.Flags&net.FlagLoopback != 0 {
			status = append(status, "LOOPBACK")
		}

		if i.Flags&net.FlagUp != 0 {
			status = append(status, "UP")
		} else {
			status = append(status, "DOWN")
		}

		if i.Flags&net.FlagRunning != 0 {
			status = append(status, "RUNNING")
		}

		inter.Status = strings.Join(status, "|")

		if addrs, err := i.Addrs(); err == nil {
			for _, addr := range addrs {
				var ipnet *net.IPNet = addr.(*net.IPNet)
				if ipnet.IP.To4() != nil {
					ipInfo := strings.Split(ipnet.String(), "/")
					inter.IPAddr = ipInfo[0]
					inter.Netmask = toIPv4(ipInfo[1])
				}
			}
		}
		interfs = append(interfs, inter)
	}
	return interfs, nil
}
