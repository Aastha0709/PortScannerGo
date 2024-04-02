package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  int
	State string
}

func ScanPort(protocol, hostname string, port int) ScanResult {

	ans := ScanResult{Port: port}

	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		ans.State = "closed"
		return ans
	}

	defer conn.Close()
	ans.State = "open"
	return ans
}

func InitialScan(hostname string) []ScanResult {

	var results []ScanResult

	for i := 1; i <= 1024; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}
	return results
}
