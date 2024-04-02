package main

import (
	"fmt"

	"PortScanner/port"
)

func main() {
	fmt.Println("Port Scanner in Go")
	result := port.ScanPort("tcp", "localhost", 443)
	fmt.Printf("Port : %t\n", result)

	results := port.InitialScan("localhost")
	fmt.Println(results)
}
