package network


import (
	"fmt"
	"net"
	"strings"
	"time"
	"os"
	"io/ioutil"
	"net/http"

	"github.com/jackpal/gateway"
)

type Connectivity struct {
	DNSResolution bool
	Internet bool
	ActiveIP string
	PublicIP string
	DefaultGateway string
}

func (c *Connectivity) String() string {
	var out string
	out += fmt.Sprintf("%-20s %-10v\n", "DNS Resolution:", c.DNSResolution)
	out += fmt.Sprintf("%-20s %-10v\n", "Internet:", c.Internet)
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

func GetConnectivityInfo() (*Connectivity, string) {
	c := &Connectivity{ActiveIP: "Unknown", PublicIP: "Unknown", DefaultGateway: "Unknown"}

	gatewayIP, err := gateway.DiscoverGateway()
	if err != nil {
		return c, fmt.Sprintln("Unable to determine default gateway:", err.Error())
	}
	c.DefaultGateway = gatewayIP.String()

	ips, err := net.LookupIP("google.com")
	if err != nil {
		return c, fmt.Sprintln("DNS resolution error:", err.Error())
	}
	c.DNSResolution = true

	conn, err := net.DialTimeout("tcp", ips[0].String() + ":443", time.Duration(3) * time.Second)
	if err != nil {
		return c, fmt.Sprintln("Internet connectivity error:", err.Error())
	}
	c.Internet = true

	c.ActiveIP = strings.Split(conn.LocalAddr().String(), ":")[0]
	c.PublicIP, err = getPublicIP()
	if err != nil {
		return c, fmt.Sprintln("Unable to determine public IP:", err.Error())
	}



	return c, ""
}

func getPublicIP() (string, error) {
	url := "https://api.ipify.org?format=text"	// we are using a pulib IP API, we're using ipify here, below are some others
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
