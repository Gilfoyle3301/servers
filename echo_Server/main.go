package main

import (
	"io"
	"log"
	"net"
	"sync"
)

func echoHandler(connection net.Conn, wg *sync.WaitGroup) {
	defer connection.Close()
	var buffer []byte = make([]byte, 1096)
	for {
		bytes, err := connection.Read(buffer[0:])

		switch {
		case err == io.EOF:
			log.Println("connection of client is close")
		case err != nil:
			log.Println("error detected ", err)
			break
		}

		log.Printf("Received %d bytes: %s\n", bytes, string(buffer))
		if _, err := connection.Write(buffer[:bytes]); err != nil {
			log.Fatalln("Unable writing data")
		}
		log.Println("Return data succes")
		wg.Done()

	}

}

func main() {
	var wg sync.WaitGroup
	dial, err := net.Listen("tcp", "[::1]:5555")
	if err != nil {
		log.Fatalln("error to bind port")
	}
	log.Println("listen addres 0.0.0.0:5555")

	connection, err := dial.Accept()
	if err != nil {
		log.Println("error connection")
	}
	wg.Add(1)
	go echoHandler(connection, &wg)
	wg.Wait()

}
