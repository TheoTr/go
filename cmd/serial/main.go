package main

import (
	goio "io"
	"log"

	"github.com/TheoTr/go/io"
)

func main() {
	badSerial := io.NewSlowWriter(3)
	err := SerialSend(badSerial, "AT+JOIN")
	if err != nil {
		log.Fatal(err)
	}

}

func SerialSend(serial goio.Writer, atCommand string) error {

	for len(atCommand) > 0 {

		n, err := serial.Write([]byte(atCommand))
		if err != nil {
			return err
		}
		atCommand = atCommand[n:]
		log.Println(n, atCommand)

	}

	return nil
}
