package port

import (
	"net"
	"strconv"
	"sync"
	"time"
)

type ScanResult struct {
	Port  string
	State string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}
	addr := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, addr, 5*time.Second)

	if err != nil {
		// log.Printf("Failed to establish network connection. %+v\n", err)
		result.State = "closed"
		return result
	}
	defer conn.Close()

	result.State = "open"
	return result
}

func InitialScan(hostname string, numOfPorts int, enableThreads bool) []ScanResult {
	var results []ScanResult

	// handle basic bad input for ports
	if numOfPorts > 65535 {
		numOfPorts = 65535
	} else if numOfPorts < 1 {
		numOfPorts = 1024
	}

	if !enableThreads {
		for i := 1; i <= numOfPorts; i++ {
			results = append(results, ScanPort("tcp", hostname, i))
			results = append(results, ScanPort("udp", hostname, i))
		}

		return results
	}

	var wg sync.WaitGroup
	var muResults sync.Mutex

	for i := 1; i <= numOfPorts; i++ {
		port := i
		wg.Go(func() {
			tcpRes := ScanPort("tcp", hostname, port)
			udpRes := ScanPort("udp", hostname, port)

			muResults.Lock()
			results = append(results, tcpRes, udpRes)
			muResults.Unlock()
		})
	}
	wg.Wait()
	return results
}
