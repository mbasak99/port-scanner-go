package port

import (
	"log"
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  int
	State string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	addr := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, addr, time.Minute)

	if err != nil {
		log.Printf("Failed to establish network connection. %+v\n", err)
		return ScanResult{Port: port, State: "closed"}
	}
	defer conn.Close()

	return ScanResult{Port: port, State: "open"}
}

func InitialScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 1; i <= 1024; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}

	return results
}
