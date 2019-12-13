package core

import (
	"csv_parser/pkg/database"
	"csv_parser/pkg/models"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ReadCsvFile(filePath string) error  {
	csvfile, err := os.Open(filePath)
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
			RegNumber:    data[0],
			Name:  data[1],
			URL:   data[2],
			Phone: data[3],
			Mail:  data[4],
		})
	}
	err = database.InsertRecords(records)
	if err != nil {
		return err
	}
	return nil
}