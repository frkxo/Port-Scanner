package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResults struct {
	Port    string
	Status  string
	Service string
}

func PortScan(protocol, hostName string, port int) ScanResults {

	results := ScanResults{
		Port: strconv.Itoa(port) + string(" : ") + protocol,
	}
	address := hostName + ":" + strconv.Itoa(port)
	connect, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		results.Status = "Port Closed"
		return results
	}
	defer connect.Close()
	results.Status = "Port Open"
	return results
}

func InitialScan(hostName string) []ScanResults {

	var result []ScanResults

	for i := 0; i <= 1024; i++ {
		result = append(result, PortScan("tcp", hostName, i))
	}

	for i := 0; i <= 1024; i++ {
		result = append(result, PortScan("udp", hostName, i))
	}

	return result
}
