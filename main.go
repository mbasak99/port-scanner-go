package main

import (
	"fmt"

	"github.com/mbasak99/tcp-port-scanner/port"
)

func main() {
	fmt.Println("Port Scanner in Go...")
	fmt.Println("Beginning scan on network.")
	ports := []int{1313, 22, 80, 443, 3001}

	// May return true based on if a process is using one of the above ports.
	for _, portNum := range ports {
		open := port.ScanPort("tcp", "localhost", portNum)
		fmt.Printf("Port state: %v | Port: %d\n", open.State, open.Port)
	}
	// Should always return true
	open := port.ScanPort("tcp", "google.com", 80)
	fmt.Printf("Port state: %v | Port: %d\n", open.State, open.Port)

	// Scan ports 1-1024
	fmt.Println("Scanning for process on ports 1-1024...")
	initScan := port.InitialScan("localhost")
	fmt.Println(initScan)
}
