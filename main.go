package main

import (
	"PortScanner/port"
	"bufio"
	"fmt"
	"log"
	"os"
)

func writeFile() error {
	fd, err := os.OpenFile(`UDP_OpenedList.txt`, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf(`os.OpenFile() -> %w`, err)
	}
	defer fd.Close()

	w := bufio.NewWriter(fd)
	for _, v := range port.UdpList {
		fmt.Fprintln(w, `UDP:`, v, `Opened!`)
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln(`Usage: ./PortScanner [ip_address]`)
	}

	log.Println("\x1b[32;4mPortScanner activited!\x1b[0m")

	//// Scan UDP port
	for i := 1; i <= 65535; i += 1 {
		port.Waitgroup.Add(1)
		go port.ScanUdp(os.Args[1], i)
	}
	fmt.Println("\x1b[32m\t>>> Done scanning UDP\x1b[0m")

	//// Scan TCP port
	for i := 1; i <= 65535; i += 1 {
		port.Waitgroup.Add(1)
		go port.ScanTcp(os.Args[1], i)
	}
	fmt.Println("\x1b[32m\t>>> Done scanning TCP\x1b[0m")
	port.Waitgroup.Wait()

	//// Write opened UDP port into a file
	if err := writeFile(); err != nil {
		log.Fatalln(`writeFile() ->`, err)
	}
	log.Println("\x1b[32;4mDone scanning!\x1b[0m")
}
