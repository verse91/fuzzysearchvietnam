# Fuzzy Search hõ trợ tiếng Việt
## Install
```
go get github.com/verse91/fuzzysearchvietnam
```
## Example
```golang
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
```
## Output
```
Nhập từ khóa: vinh 
Kết quả tìm được:
59. Trà Vinh
62. Vĩnh Phúc
61. Vĩnh Long
64. va in gh
```
