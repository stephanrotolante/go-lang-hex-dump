package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

var err error

func main() {
	var inputFilePath string

	flag.StringVar(&inputFilePath, "f", "", "path to input file")
	flag.Parse()

	fileInfo, err := os.Stat(inputFilePath)
	if err != nil {
		fmt.Printf("Error occured getting file stats %s", inputFilePath)
		log.Panic(err)
	}

	if fileInfo.IsDir() {
		log.Panic(errors.New("cannot hex dump directory"))
	}

	// Get the size of the file
	fileSize := fileInfo.Size()

	kilobytes := float64(fileSize) / 1024

	megabytes := kilobytes / 1024

	gigbytes := megabytes / 1024
	fmt.Printf("File size:\n%d bytes\n%f kilobytes\n%f megabytes\n%f gigabytes\n\n", fileSize, kilobytes, megabytes, gigbytes)

	file, err := os.Open(inputFilePath)

	defer file.Close()

	if err != nil {
		fmt.Printf("Error occured opening file %s", inputFilePath)
		log.Panic(err)
	}

	var counter int = 0

	buffer := make([]byte, 16)

	for {

		buffer = []byte{
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
		}

		n, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error:", err)
			break
		}
		if n == 0 {
			break
		}

		fmt.Printf("%08x: ", counter)

		for i := 0; i < n; i++ {
			fmt.Printf("%02x ", buffer[i])
		}
		counter++

		fmt.Println("")
	}

}
