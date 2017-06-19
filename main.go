package main

import (
	"flag"
	"fmt"
    "log"
	"net"
	"strconv"
	"strings"
)

const (
	stopCharacter = "\r\n\r\n"
	exitCommand = "exit"
)

func socketClient(ip string, port int, stat string) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	cmd := fmt.Sprintf("metrics-single:%s", stat)
	_, err = conn.Write([]byte(cmd))
	if err != nil {
		log.Fatalln(err)
	}
	_, err = conn.Write([]byte(stopCharacter))
	if err != nil {
		log.Fatalln(err)
	}
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(string(buff[:n]))
	conn.Write([]byte(exitCommand))
}

func main() {

	ip := flag.String("ip", "127.0.0.1", "the ip address of the perfmon agent")
	port := flag.Int("port", 4444, "the port that the perfmon agent is listening")
	stat := flag.String("stat", "memory", "the stat to read")
	flag.Parse()
	socketClient(*ip, *port, *stat)
}