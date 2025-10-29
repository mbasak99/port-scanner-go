package main

import (
	"fmt"
	"time"

	"github.com/mbasak99/tcp-port-scanner/port"
)

func main() {
	fmt.Println("Port Scanner in Go...")
	fmt.Println("Beginning scan on network.")
	// ports := []int{1313, 22, 80, 443, 3001}

	// // May return true based on if a process is using one of the above ports.
	// for _, portNum := range ports {
	// 	open := port.ScanPort("tcp", "localhost", portNum)
	// 	fmt.Printf("Port state: %v | Port: %s\n", open.State, open.Port)
	// }
	// // Should always return true
	// open := port.ScanPort("tcp", "google.com", 80)
	// fmt.Printf("Port state: %v | Port: %s\n", open.State, open.Port)

	// Scan ports 1-1024
	// Sync
	numPorts := 1024
	fmt.Println("Scanning for process on ports 1-1024...")
	start := time.Now()
	/*initScan :=*/ port.InitialScan("localhost", numPorts, false)
	stop := time.Since(start)
	// fmt.Println(initScan)

	// Async
	startAsync := time.Now()
	/*initScan =*/ port.InitialScan("localhost", numPorts, true)
	stopAsync := time.Since(startAsync)
	fmt.Printf("Time taken sync: %s\n", stop)
	fmt.Printf("Time taken async: %s\n", stopAsync)
	fmt.Printf("Go threads delta: %.2fx\n", float64(stopAsync.Milliseconds())/float64(stop.Milliseconds()))

	// Scan ports 1-65535
	// Sync
	numPorts = 65535
	fmt.Println("\nScanning for process on ports 1-65535...")
	start = time.Now()
	port.InitialScan("localhost", numPorts, false)
	stop = time.Since(start)

	// Async
	startAsync = time.Now()
	/*initScan =*/ port.InitialScan("localhost", numPorts, true)
	stopAsync = time.Since(startAsync)
	fmt.Printf("Time taken sync: %s\n", stop)
	fmt.Printf("Time taken async: %s\n", stopAsync)
	fmt.Printf("Go threads delta: %.2fx\n", float64(stopAsync.Milliseconds())/float64(stop.Milliseconds()))
}
