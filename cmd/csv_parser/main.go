package main

import (
	"csv_parser/pkg/core"
	"log"
)

var (
	csvfile = "./asset.csv"
)

func main() {
	err := core.ReadCsvFile(csvfile)
	if err != nil {
		log.Fatal("failed to parse csv document")
	}
}
