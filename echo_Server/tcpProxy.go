package main

import (
	"io"
	"log"
	"net"
)

func handleProxy(conn net.Conn) {
	dst, ok := net.Dial("tcp", "77.88.55.88:80")
	if ok != nil {
		log.Fatalln(ok)
	}
	defer dst.Close()
	go func() {
		if _, err := io.Copy(dst, conn); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := io.Copy(conn, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	src, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := src.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleProxy(conn)
	}

}
