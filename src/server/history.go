package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const HistoryTable = "history_table"

const WaitedGuest = "Stranger"

const timeZone = "Local"
const timeFormat = "2006-01-02 03:04:05"

type Record struct {
	Id            int64
	Name          string
	ImageLocation string
	Status        string
	Time          string
}

//.
//.
//.
//.
//All the grpc functions for the app to server for the history record part
func (rs *routeServer) GetHistoryRecorded(ctx context.Context, timestamp *pb.Timestamp) (*pb.HistoryRecords, error) {
	var err error
	var recordList pb.HistoryRecords
	records, err := rs.GetHisRecDBbyTime(timestamp.Starttime, timestamp.Endtime)
	if err != nil {
		return &recordList, err
	}

	//loop through every records in the records list
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

//get history image from the ImageStore
func (rs *routeServer) GetHistoryImage(imageuuid *pb.ImageLocation, stream pb.Route_GetHistoryImageServer) error {
	f, err := os.Open(rs.imagestore + "/" + imageuuid.Address)
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
	fmt.Println("Sent %d bytes", sizeTotal)
	return nil
}

func (rs *routeServer) GivePermission(ctx context.Context, permission *pb.Permission) (*pb.Empty, error) {

	var err error
	if permission.Usernames != WaitedGuest {
		return nil, status.Errorf(codes.NotFound, "Permission Guest Name did not match!")
	}

	//set the getPermission map with id , value according to the permission

	//update the permission status in the database
	if permission.Permit {
		rs.waitingUser <- permAllow
	} else {
		rs.waitingUser <- permDeny
	}

	return nil, err
}

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

//.
//.
//.
//.
//All the functions for the communication between server to database for history record part.

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

func (rs *routeServer) UpdateRecordStatusToDB(id int64, status string) error {
	sql_q := fmt.Sprintf(
		"UPDATE `%s` SET status = '%s' where id = '%d';",
		HistoryTable, status, id)
	_, err := rs.conn.Exec(sql_q)

	return err
}

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

func (rs *routeServer) ClearHistoryRecord() error {
	sql_q := fmt.Sprintf(
		"DELETE FROM `%s` WHERE id > 0;", HistoryTable)
	_, err := rs.conn.Exec(sql_q)
	return err
}

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
