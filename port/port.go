package port

import (
	"log"
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  string
	State string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}
	addr := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, addr, time.Minute)

	if err != nil {
		log.Printf("Failed to establish network connection. %+v\n", err)
		result.State = "closed"
		return result
	}
	defer conn.Close()

	result.State = "open"
	return result
}

func InitialScan(hostname string, numOfPorts int) []ScanResult {
	var results []ScanResult

	// handle basic bad input for ports
	if numOfPorts > 65535 {
		numOfPorts = 65535
	} else if numOfPorts < 1 {
		numOfPorts = 1024
	}

	for i := 1; i <= numOfPorts; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
		results = append(results, ScanPort("udp", hostname, i))
	}

	return results
}
