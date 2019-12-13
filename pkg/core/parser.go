package core

import (
	"csv_parser/pkg/database"
	"csv_parser/pkg/models"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadCsvFile(filePath string) (err error) {
	// iterate over directories
	searchDir := "./import"

	fileList := []string{}
	err = filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	// TODO: implement concurently run
	for _, file := range fileList {
		if strings.Contains(file, ".csv") {
			fmt.Println(file)
			csvfile, err := os.Open(file)
			if err != nil {
				log.Fatalln("Couldn't open the csv file", err)
			}

			r := csv.NewReader(csvfile)
			r.Comma = '|'
			var records []models.Record
			for {
				data, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}

				records = append(records, models.Record{
					RegNumber: data[0],
					Name:      data[1],
					URL:       data[2],
					Phone:     data[3],
					Mail:      data[4],
				})
			}
			err = database.InsertRecords(records)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
