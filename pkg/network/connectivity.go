package network

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jackpal/gateway"
)

type Connectivity struct {
	DNSResolution  bool
	Internet       bool
	ActiveInterface string
	ActiveIP       string
	PublicIP       string
	DefaultGateway string
}

func (c *Connectivity) String() string {
	var out string
	out += fmt.Sprintf("%-20s %-10v\n", "DNS Resolution:", c.DNSResolution)
	out += fmt.Sprintf("%-20s %-10v\n", "Internet:", c.Internet)
	out += fmt.Sprintf("%-20s %-10v\n", "Active Interface:", c.ActiveInterface)
	out += fmt.Sprintf("%-20s %-10s\n", "Active Local IP:", c.ActiveIP)
	out += fmt.Sprintf("%-20s %-10s\n", "Public IP:", c.PublicIP)
	out += fmt.Sprintf("%-20s %-10s\n", "Default Gateway:", c.DefaultGateway)

	return out
}

func GetHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return hostname, nil
}

func GetConnectivityInfo(interfaces []*CustomInterface) (*Connectivity, error) {
	c := &Connectivity{ActiveIP: "Unknown", PublicIP: "Unknown", DefaultGateway: "Unknown"}

	gatewayIP, err := gateway.DiscoverGateway()
	if err != nil {
		return c, fmt.Errorf("Unable to determine default gateway: %v\n", err)
	}
	c.DefaultGateway = gatewayIP.String()

	ips, err := net.LookupIP("google.com")
	if err != nil {
		return c, fmt.Errorf("DNS resolution error: %v\n", err)
	}
	c.DNSResolution = true

	conn, err := net.DialTimeout("tcp", ips[0].String()+":443", time.Duration(3)*time.Second)
	if err != nil {
		return c, fmt.Errorf("Internet connectivity error: %v\n", err)
	}
	c.Internet = true

	c.ActiveIP = strings.Split(conn.LocalAddr().String(), ":")[0]
	for _, interf := range interfaces {
		if interf.IPAddr == c.ActiveIP {
			c.ActiveInterface = interf.Name
			break
		}
	}
	c.PublicIP, err = getPublicIP()
	if err != nil {
		return c, fmt.Errorf("Unable to determine public IP: %v\n", err)
	}

	return c, nil
}

func getPublicIP() (string, error) {
	url := "https://api.ipify.org?format=text" // we are using a pulib IP API, we're using ipify here, below are some others
	// https://www.ipify.org
	// http://myexternalip.com
	// http://api.ident.me
	// http://whatismyipaddress.com/api
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}
