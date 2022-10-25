package main

import (
	"log"

	"github.com/TheoTr/go/io"
)

func main() {
	atCommand := "AT+JOIN"
	badSerial := io.NewSlowWriter(3)

	for len(atCommand) > 0 {

		n, err := badSerial.Write([]byte(atCommand))
		if err != nil {
			log.Fatal(err)
		}
		atCommand = atCommand[n:]
		log.Println(n, atCommand)

	}

}
