package main

import (
	"bufio"
	"context"
	"fmt"
	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"github.com/CPEN391-Team-4/backend/src/logging"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Table for history storage
const HistoryTable = "history_table"

// Time zone to be used in the database
const timeZone = "Local"

// Format for date stamps stored in the database
const timeFormat = "2006-01-02 03:04:05"

// Representation of a history record as stored in the database
type Record struct {
	Id            int64
	Name          string
	ImageLocation string
	Status        string
	Time          string
}

/********************* DATABASE *********************/

// AddRecordToDB Add a history record entry into HistoryTable
func (rs *routeServer) AddRecordToDB(name string, imageLocation string) (int64, error) {

	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return 0, err
	}

	dt := time.Now().In(loc)
	recordTime := dt.Format(timeFormat)

	sql_q := fmt.Sprintf(
		"INSERT INTO `%s` (name, status, ImageLocation, time) VALUES ('%s', '%s', '%s', '%s');",
		HistoryTable, name, "unknown", imageLocation, recordTime)
	res, err := rs.conn.Exec(sql_q)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := res.LastInsertId()

	return lastInsertId, err
}

// UpdateRecordStatusToDB Update the status for record with id
func (rs *routeServer) UpdateRecordStatusToDB(id int64, status string) error {
	sql_q := fmt.Sprintf(
		"UPDATE `%s` SET status = '%s' where id = '%d';",
		HistoryTable, status, id)
	_, err := rs.conn.Exec(sql_q)

	return err
}

// GetHisRecDBbyTime Get all records between starttime and endtime
func (rs *routeServer) GetHisRecDBbyTime(starttime string, endtime string) ([]Record, error) {
	sql_q := "SELECT * FROM " + HistoryTable + " WHERE time between ? and ?"
	res, err := rs.conn.Query(sql_q, starttime, endtime)

	if err != nil {
		return nil, err
	}

	var records []Record

	for res.Next() {
		var record Record
		err = res.Scan(&record.Id, &record.Name, &record.Status, &record.ImageLocation, &record.Time)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, record)
	}

	return records, err
}

// DeleteRecordFromDB Delete the record of the image at imagelocation
func (rs *routeServer) DeleteRecordFromDB(imagelocation string) error {
	sql_q := fmt.Sprintf(
		"DELETE FROM `%s` WHERE  ImageLocation ='%s';", HistoryTable, imagelocation)
	_, err := rs.conn.Exec(sql_q)
	return err
}

// ClearHistoryRecord Delete all from HistoryTable
func (rs *routeServer) ClearHistoryRecord() error {
	sql_q := fmt.Sprintf(
		"DELETE FROM `%s` WHERE id > 0;", HistoryTable)
	_, err := rs.conn.Exec(sql_q)
	return err
}

// GetLatestRecordImageID Get the id of the latest image in HistoryTable
func (rs *routeServer) GetLatestRecordImageID() (string, error) {
	res, err := rs.conn.Query("SELECT ImageLocation FROM " + HistoryTable + " ORDER BY id DESC LIMIT 1")
	imageid := ""
	if err != nil {
		return "", err
	}
	for res.Next() {
		err = res.Scan(&imageid)
	}

	return imageid, err

}

/********************* HISTORY API *********************/

// GetHistoryRecorded Retrieve all history between timestamp.Starttime and timestamp.Endtime
func (rs *routeServer) GetHistoryRecorded(ctx context.Context, timestamp *pb.Timestamp) (*pb.HistoryRecords, error) {
	var err error
	var recordList pb.HistoryRecords
	records, err := rs.GetHisRecDBbyTime(timestamp.Starttime, timestamp.Endtime)
	if err != nil {
		return &recordList, err
	}

	for _, rec := range records {
		var record pb.HistoryRecord
		record.Name = rec.Name
		record.Status = rec.Status
		record.ImageLocation = rec.ImageLocation
		record.Time = rec.Time

		recordList.Record = append(recordList.Record, &record)
	}
	return &recordList, err
}

//GetHistoryImage Retrieve a history image from the ImageStore
func (rs *routeServer) GetHistoryImage(imageuuid *pb.ImageLocation, stream pb.Route_GetHistoryImageServer) error {
	file := rs.imagestore + "/" + imageuuid.Address
	f, err := os.Open(file)
	log.Printf("Requested file=%s", file)
	if err != nil {
		return logging.LogError(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, READ_BUF_SIZE)

	var photo pb.Photo

	sizeTotal := 0
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		photo.Image = buf[0:n]
		err = stream.Send(&photo)
		if err != nil {
			return err
		}
		sizeTotal += n
	}

	return nil
}

// DeleteRecords Delete the record at imageid.Address
func (rs *routeServer) DeleteRecords(ctx context.Context, imageid *pb.ImageLocation) (*pb.Empty, error) {
	var err error
	err = rs.DeleteRecordFromDB(imageid.Address)
	if err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, err
}

// DeleteRecords Delete the record at imageid.Address
func (rs *routeServer) GetLatestImage(_ *pb.Empty, stream pb.Route_GetLatestImageServer) error {
	var err error
	imageid, err := rs.GetLatestRecordImageID()

	f, err := os.Open(rs.imagestore + "/" + imageid)
	if err != nil {
		return err
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, READ_BUF_SIZE)

	var photo pb.Photo

	sizeTotal := 0
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		photo.Image = buf[0:n]
		err = stream.Send(&photo)
		if err != nil {
			return err
		}
		sizeTotal += n
	}
	log.Printf("The latest image id is %s\n", imageid)
	log.Printf("Sent %d bytes\n", sizeTotal)

	return err
}