package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args
	argsWithoutProg := args[1:]

	if len(argsWithoutProg) == 0 {
		log.Fatal("File argument is required")

	}

	fileName := argsWithoutProg[0]

	with_os_open(fileName)
	with_os_open_io_copy(fileName)

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)

	}

	value := string(file)

	fmt.Println(value)
}

func with_os_open(fileName string) {

	file, error := os.Open(fileName)

	if error != nil {
		log.Fatalf("unable to open the file: %v", error)
		return
	}

	defer file.Close()

	if error != nil {
		log.Fatalf("Error reading file: %v", error)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file")
	}

}
func with_os_open_io_copy(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
		return
	}
	defer file.Close()

	bytesWritten, err := io.Copy(os.Stdout, file)
	if err != nil {
		log.Fatalf("Error copying file: %v", err)
		return
	}

	fmt.Printf("Copied %d bytes from %s to standard output.\n", bytesWritten, fileName)
}

/*

Feature                      os.Open                            os.ReadFile
Operation                    Opens file for reading/writing       Reads entire file into memory
Control                      More control over file operations    Simpler, less control
Resource Management          Requires manual file closing         Handles file closing automatically
Use Cases                    Streaming, chunk reading, complex file operations   Reading small/medium files entirely
Memory Usage                 Lower for large files (when reading in chunks)   Can be high for large files

*/
