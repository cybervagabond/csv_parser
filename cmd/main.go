package main

import (
	"csv_parser/pkg/core"
	"log"
)

func main() {
	err := core.ReadCsvFile("./asset.csv")
	if err != nil {
		log.Fatal("failed to parse csv document")
	}
}

