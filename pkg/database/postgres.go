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
		// check that record exists
		queryCheck := fmt.Sprintf(`select exists(select reg_number from records where reg_number='%s');`,
			record.RegNumber)

		rows, err := conn.Query(queryCheck)
		if err != nil {
			return err
		}
		var exists bool
		for rows.Next() {
			err = rows.Scan(&exists)
			if err != nil {
				return err
			}
		}
		if exists == false {
			log.Println("creating new record: ", record)
			queryInsert := fmt.Sprintf(`INSERT INTO records(reg_number, name, url, phone, mail)
 VALUES ('%s', '%s', '%s', '%s', '%s');`,
				record.RegNumber,
				record.Name,
				record.URL,
				record.Phone,
				record.Mail)

			_, err = conn.Exec(queryInsert)
			if err != nil {
				return err
			}
		} else {
			log.Printf("record with reg_number %s already exists", record.RegNumber)
		}
	}
	return nil
}
