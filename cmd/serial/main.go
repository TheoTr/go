package main

import (
	"log"

	"github.com/TheoTr/go/io"
)

func main() {
	atCommand := "AT+JOIN"
	badSerial := io.NewSlowWriter(3)

	log.Println(atCommand)
	n, err := badSerial.Write([]byte(atCommand))
	if err != nil {
		log.Fatal(err)
	}

	atCommand = atCommand[3:]
	log.Println(n, atCommand)

	n, err = badSerial.Write([]byte(atCommand))
	if err != nil {
		log.Fatal(err)
	}

	atCommand = atCommand[3:]
	log.Println(n, atCommand)

	n, err = badSerial.Write([]byte(atCommand))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(n, atCommand)
}
