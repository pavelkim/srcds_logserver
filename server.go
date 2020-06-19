package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"plugin"
	"syscall"
)

var ApplicationDescription string = "UDP Server"
var BuildVersion string = "0.0.0a"
var Debug bool = false

func handleSignal() {
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChannel
		log.Print("SIGINT")
		os.Exit(0)
	}()
}

func handlePacket(data []byte) {
	log.Print("Handling a packet length=", len(data))
}

func main() {

	bindPtr := flag.String("bind", "127.0.0.1:9001", "Address and port to listen")
	handlerPtr := flag.String("handler", "handler.so", "Path to the payload handler shared library")
	enableDebugPtr := flag.Bool("debug", false, "Enable verbose output")
	showVersionPtr := flag.Bool("version", false, "Show version")

	flag.Parse()

	if *enableDebugPtr {
		Debug = true
	}

	if *showVersionPtr {
		fmt.Printf("%s\n", ApplicationDescription)
		fmt.Printf("Version: %s\n", BuildVersion)
		os.Exit(0)
	}

	payloadHandlerFilename, err := plugin.Open(*handlerPtr)
	if err != nil {
		log.Fatal("Error while opening plugin file:", err)
	}

	symbol, err := payloadHandlerFilename.Lookup("PayloadHandlerFunction")
	if err != nil {
		log.Fatal("Error while looking up a symbol:", err)
	}

	payloadHandler := symbol.(func([]byte) (bool, error))

	listen_address, err := net.ResolveUDPAddr("udp4", *bindPtr)
	if err != nil {
		log.Fatal("Error while resolving address: ", err)
	}

	connection, err := net.ListenUDP("udp", listen_address)
	if err != nil {
		log.Fatal("Error while opening socket: ", err)
	}

	defer connection.Close()

	log.Print("Listening on ", listen_address.String())

	read_buffer := make([]byte, 65535)

	for {

		length, remote, err := connection.ReadFrom(read_buffer)
		if err != nil {
			panic(err)
		}

		go func() {

			log.Print("Accepted UDP packet from ", remote, " length=", length)
			payload := read_buffer[:length]

			if Debug {
				log.Printf("Dump:\n%s", hex.Dump(payload))
			}

			if len(payload) < 1 {
				log.Printf("Warning: Packet is too short!\n")
			}

			payloadHandler(payload)

		}()
	}

}
