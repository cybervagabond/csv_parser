package database

import (
	"fmt"
	"log"
	"time"
)

func writeLog(logMsg string) {
	conn := connectDb("records app")
	defer conn.Close()

	query := fmt.Sprintf(`INSERT INTO records_log(time, event) VALUES('%s','%s')`,
		time.Now().Format(time.RFC3339), logMsg)
	_, err := conn.Query(query)
	if err != nil {
		log.Println(err)
	}
}
