package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const History_TABLE = "history_table"

type Record struct {
	Id             int64
	Name           string
	Image_location string
	Status         string
	Time           string
}

func addRecordToDB(db *sql.DB, name string, image_location string) (int64, error) {

	loc, _ := time.LoadLocation("MST")
	dt := time.Now().In(loc)
	recordtime := dt.Format(time.RFC3339)

	sql := fmt.Sprintf(
		"INSERT INTO `%s` (name, status, ImageLocation, time)VALUES ('%s', '%s', '%s', '%s');",
		History_TABLE, name, "unknown", image_location, recordtime)
	res, err := db.Exec(sql)

	last_insert_id, err := res.LastInsertId()
	//fmt.Println(last_insert_id)

	return last_insert_id, err
}

func updateRecordStatusToDB(db *sql.DB, id int64, status string) error {
	sql := fmt.Sprintf(
		"UPDATE `%s` SET status = '%s' where id = '%d';",
		History_TABLE, status, id)
	_, err := db.Exec(sql)

	return err
}

func getHisRecDBbyTime(db *sql.DB, starttime string, endtime string) ([]Record, error) {

	res, err := db.Query("SELECT * FROM history_table where time between ? and ?", starttime, endtime)

	if err != nil {
		return nil, err
	}

	var records []Record

	for res.Next() {
		var record Record
		fmt.Println("onerow")
		err = res.Scan(&record.Id, &record.Name, &record.Status, &record.Image_location, &record.Time)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, record)
	}

	return records, err
}

func clearHistoryRecord(db *sql.DB) error {
	sql := fmt.Sprintf(
		"Delete from `%s` where id > 0;", History_TABLE)
	_, err := db.Exec(sql)
	return err
}

func main() {
	fmt.Println("entertesthistory")
	// environ := env{}
	// environ.readEnv()
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/cpen391_backend")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()
	id1, err := addRecordToDB(db, "stranger", "imagelocation1")
	updateRecordStatusToDB(db, id1, "reject")
	id2, err := addRecordToDB(db, "stranger", "imagelocation2")
	updateRecordStatusToDB(db, id2, "reject")
	id3, err := addRecordToDB(db, "stranger", "imagelocation3")
	updateRecordStatusToDB(db, id3, "reject")
	id4, err := addRecordToDB(db, "stranger", "imagelocation4")
	updateRecordStatusToDB(db, id4, "reject")

	// records, err := getHisRecDBbyTime(db, "2021-03-25 12:00:00", "2021-03-25 13:50:00")
	// fmt.Println(records)
	// fmt.Println(err)

	//clearHistoryRecord(db)

}
