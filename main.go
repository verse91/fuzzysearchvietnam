package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/verse91/fuzzysearchvietnam"
)

func main() {
	// Đọc file list.txt
	file, err := os.Open("list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var items []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Nhập từ khóa
	var query string
	fmt.Print("Nhập từ khóa: ")
	fmt.Scanln(&query)

	// Tìm kiếm
	results := fuzzyvn.Find(query, items)

	// In kết quả
	fmt.Println("Kết quả tìm được:")
	for _, res := range results {
		fmt.Println(res)
	}
}
