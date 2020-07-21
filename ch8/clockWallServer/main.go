// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 219.
//!+

// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// go run main.go Asia/Shanghai=8000 US/Eastern=8001
func main() {
	args := os.Args[1:]
	for _, arg := range args {
		params := strings.Split(arg, "=")
		timeZone, port := params[0], params[1]
		listener, err := net.Listen("tcp", "localhost:"+port)
		if err != nil {
			log.Fatal(err)
		}
		go handListener(listener, timeZone)
	}
	for {
		time.Sleep(time.Second)
	}
}

func handListener(listener net.Listener, timeZone string) {
	for {
		conn, err := listener.Accept()	//这是一个阻塞，直到收到连接请求才会有返回并进行下一步
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn, timeZone) // handle one connection at a time
	}
}

func handleConn(c net.Conn, timeZone string) {
	defer c.Close()
	var cstSh, _ = time.LoadLocation(timeZone)
	for {
		_, err := io.WriteString(c, time.Now().In(cstSh).Format("15:04:05\n"))	//net.Conn 满足 io.Writer 接口，可以直接向其写入
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

//!-
