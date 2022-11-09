package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func createFile() {
	// example of opening a file in the read-only mode
	file, err := os.Create("new_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// closing the file
	file.Close()
}

func removeFile() {
	err := os.Remove("new_file.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func removeAll() {
	err := os.RemoveAll("test_directory")
	if err != nil {
		log.Fatal(err)
	}
}

func readFile() {
	data, err := os.ReadFile("test_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func readFilePerLine() {
	file, err := os.Open("test_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readFilePerWord() {
	file, err := os.Open("song.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords) // split each scanned line into words

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readFilePerChunk() {
	const chunkSize = 8

	file, err := os.Open("test_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, chunkSize)
	for {
		readTotal, err := file.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(string(buf[:readTotal]))
	}
}

func writingFiles() {
	file, err := os.Create("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := []string{"Line 1", "Line 2", "Line 3"}
	for i, line := range data {
		fmt.Fprintln(file, line)
		if err != nil {
			log.Fatal(err)
		}
		if i == len(data)-1 {
			fmt.Printf("%d lines written successfully!", i+1)
		}
	}
}

func appendingData() {
	// open the file in append-and-write mode - permission mode 0644 is required!
	file, err := os.OpenFile("hello.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	additionalLine := "ALWAYS 🕓 make sure 👍 your code 💻 is 💡 clean 🧼 and well 💯 structured 🏛."

	b, err := fmt.Fprintln(file, additionalLine) // append the additional line
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes appended successfully!", b) // 92 bytes appended successfully!
}

func main() {
	appendingData()
}
