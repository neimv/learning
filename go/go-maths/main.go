package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fd, err := os.Open("titanic.csv")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("File open")
	defer fd.Close()

	fileReader := csv.NewReader(fd)
	records, err := fileReader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(records[1])

	for i, s := range records[1] {
		fmt.Println(i, s)
		fmt.Printf("%T\n", s)
	}
}
