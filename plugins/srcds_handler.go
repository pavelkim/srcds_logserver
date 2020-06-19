package main

import (
	"log"
)

func PayloadHandlerFunction(payload []byte) (bool, error) {
	log.Printf("Handling payload length=%d\n", len(payload))
	return true, nil
}
