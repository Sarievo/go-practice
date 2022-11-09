package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if len(os.Args) != 2 {
		fmt.Println("Directory is not specified")
		return
	}

	var format, option string
	fmt.Println("Enter file format:")
	scanner.Scan()
	format = scanner.Text()

	fmt.Println("\nSize sorting options:\n1. Descending\n2. Ascending")
	for option != "1" && option != "2" {
		fmt.Println("\nEnter a sorting option")
		scanner.Scan()
		option = scanner.Text()

		if option != "1" && option != "2" {
			fmt.Println("\nWrong option")
		}
	}
	fmt.Println()

	files := map[int64][]string {}

	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			// fmt.Println(path)
			fileInfo, err := os.Stat(path)
			if err != nil {
				log.Fatal(err)
			}
			name := fileInfo.Name()
			if format == "" {
				// 0
				// fmt.Println(name)
				files[fileInfo.Size()] = append(files[fileInfo.Size()], path)
			} else {
				index := strings.Index(name, "." + format)
				if index != -1 {
					// 1
					// fmt.Println(name)
					files[fileInfo.Size()] = append(files[fileInfo.Size()], path)
				}
			}
		}
		return nil
	})

	entries := []int64{}
	for k, _ := range files {
		// fmt.Println(k, v)
		entries = append(entries, k)
	}

	if option == "1" {
		sort.Slice(entries, func(i, j int) bool {
			return entries[i] > entries[j]
		})
	} else {
		sort.Slice(entries, func(i, j int) bool {
			return entries[i] < entries[j]
		})
	}

	for _, x := range entries {
		fmt.Println(x, "bytes")
		for _, y := range files[x] {
			fmt.Println(y)
		}
		fmt.Println()
	}


	if err != nil {
		log.Fatal(err)
	}
}
