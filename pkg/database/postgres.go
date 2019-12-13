package database

import (
	"csv_parser/pkg/models"
	"fmt"
	"log"
)

func InsertRecords(records []models.Record) error {
	conn := connectDb("records app")
	defer conn.Close()
	for _, record := range records[1:] {
		query := fmt.Sprintf(`INSERT INTO records(reg_number, name, url, phone, mail)
 VALUES ('%s', '%s', '%s', '%s', '%s');`,
			record.RegNumber,
			record.Name,
			record.URL,
			record.Phone,
			record.Mail)

		log.Println(query)

		_, err := conn.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}
