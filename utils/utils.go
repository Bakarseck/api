package utils

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func FindAvailablePort(start int) int {
	for port:=start;port<=9999;port++ {
		address := fmt.Sprintf("localhost:%d", port)
		listener, err := net.Listen("tcp", address)
		if err == nil {
			listener.Close()
			return port
		} else {
			log.Printf("Port %d Occupé", port)
		}
	}
	return 0
}

func LoadEnv(path string) error {
	file, err := os.Open(path)
	CheckError(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		key:=strings.TrimSpace(parts[0])
		value:=strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}
	return scanner.Err()
}
