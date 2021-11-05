package port

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var Waitgroup sync.WaitGroup
var UdpList []int

// Return ture if port is opened;
// otherwise false;
func ScanPort(protocol, hostname string, port int) bool {
	address := hostname + `:` + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}

func ScanTcp(hostname string, port int) {
	defer Waitgroup.Done()

	if ScanPort("tcp", hostname, port) {
		fmt.Printf("\x1b[33mtcp/%d: Open!\x1b[0m\n", port)
	}
}

func ScanUdp(hostname string, port int) {
	defer Waitgroup.Done()

	if ScanPort("udp", hostname, port) {
		UdpList = append(UdpList, port)
		return
	}
}
