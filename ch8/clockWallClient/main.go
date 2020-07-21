// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

//go run main.go Asia/Shanghai=8000
//go run main.go US/Eastern=8001
func main() {
	args := os.Args[1:]
	for _, arg := range args {
		clock := strings.Split(arg, "=")
		_, port := clock[0], clock[1]
		conn, err := net.Dial("tcp", "localhost:"+port)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn)
	}
	for {
		time.Sleep(time.Second)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//!-
