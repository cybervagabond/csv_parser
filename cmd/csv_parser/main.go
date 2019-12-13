package main

import (
	"csv_parser/pkg/core"
	"log"
	"time"
)

var (
	csvfile = "./asset.csv"
)

func main() {
	start := time.Now()
	err := core.ReadCsvFile(csvfile)
	if err != nil {
		log.Fatal("failed to parse csv document")
	}
	log.Println("executed in: ", time.Since(start))
}
