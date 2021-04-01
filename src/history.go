package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const History_TABLE = "history_table"

const WaitedGuest = "Stranger"

type Record struct {
	Id             int64
	Name           string
	Image_location string
	Status         string
	Time           string
}

// //.
// //.
// //.
// //.
// //All the grpc function for the app to server for the history record part

// func (rs *routeServer) GetHistoryReocorded(timestamp *pb.Timestamp, stream pb.Route_GetHistoryReocordedServer) error {
// 	var err error

// 	records, err := getHisRecDBbyTime(rs.conn, timestamp.Starttime, timestamp.Endtime)
// 	if err != nil {
// 		return err
// 	}

// 	//loop through every records in the records list
// 	for i, rec := range records {

// 		// the location of this may change
// 		f, err := os.Open(rs.imagestore + "/" + rec.Image_location)
// 		if err != nil {
// 			return err
// 		}
// 		reader := bufio.NewReader(f)
// 		buf := make([]byte, READ_BUF_SIZE)
// 		var photo pb.Photo
// 		sizeTotal := 0
// 		segNumber := 0
// 		for {
// 			var record pb.HistoryRecord
// 			//only send the basic info in the first history segments
// 			if segNumber == 0 {
// 				record.Name = rec.Name
// 				record.Status = rec.Status
// 				record.Time = rec.Time
// 			}

// 			n, err := reader.Read(buf)
// 			if err != nil {
// 				if err != io.EOF {
// 					return err
// 				}
// 				break
// 			}

// 			photo.Image = buf[0:n]

// 			record.Photo = &photo
// 			err = stream.Send(&record)
// 			if err != nil {
// 				return err
// 			}
// 			sizeTotal += n
// 			segNumber += 1
// 		}
// 		log.Printf("Sent %d bytes for the %d history reocord\n", sizeTotal, i)

// 	}

// 	return err
// }

// func (rs *routeServer) GivePermission(ctx context.Context, permission *pb.Permission) (*pb.Empty, error) {

// 	var err error
// 	if permission.Usernames != WaitedGuest {
// 		return nil, status.Errorf(codes.NotFound, "Permission Guest Name did not mathch!")
// 	}

// 	//set the getPermission map with id , value according to the permission

// 	//update the permission status in the database
// 	if permission.Permit {
// 		err = updateRecordStatusToDB(rs.conn, permission.Userid, "Access")
// 	} else {
// 		err = updateRecordStatusToDB(rs.conn, permission.Userid, "Denied")
// 	}

// 	return nil, err
// }

// func (rs *routeServer) GetLastestImage(e *pb.Empty, stream pb.Route_GetLastestImageServer) error {
// 	var err error
// 	imageid, err := getLastedRecordImageID(rs.conn)

// 	f, err := os.Open(rs.imagestore + "/" + imageid)
// 	if err != nil {
// 		return err
// 	}

// 	defer f.Close()

// 	reader := bufio.NewReader(f)
// 	buf := make([]byte, READ_BUF_SIZE)

// 	var photo pb.Photo

// 	sizeTotal := 0
// 	for {
// 		n, err := reader.Read(buf)
// 		if err != nil {
// 			if err != io.EOF {
// 				return err
// 			}
// 			break
// 		}

// 		photo.Image = buf[0:n]
// 		err = stream.Send(&photo)
// 		if err != nil {
// 			return err
// 		}
// 		sizeTotal += n
// 	}
// 	log.Printf("The latest image id is %s\n", imageid)
// 	log.Printf("Sent %d bytes\n", sizeTotal)

// 	return err
// }

//.
//.
//.
//.
//All the functions for the communication between server to database for history record part.

func addRecordToDB(db *sql.DB, name string, image_location string) (int64, error) {

	loc, _ := time.LoadLocation("MST")
	dt := time.Now().In(loc)
	recordtime := dt.Format(time.RFC3339)
	fmt.Println(recordtime)

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

func getLastedRecordImageID(db *sql.DB) (string, error) {
	res, err := db.Query("select ImageLocation from history_table order by id desc limit 1")
	imageid := ""
	if err != nil {
		return "", err
	}
	for res.Next() {
		err = res.Scan(&imageid)
	}

	return imageid, err

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
	clearHistoryRecord(db)
	id1, err := addRecordToDB(db, "stranger", "imagelocation1")
	updateRecordStatusToDB(db, id1, "reject")
	id2, err := addRecordToDB(db, "stranger", "imagelocation2")
	updateRecordStatusToDB(db, id2, "reject")
	id3, err := addRecordToDB(db, "stranger", "imagelocation3")
	updateRecordStatusToDB(db, id3, "reject")
	id4, err := addRecordToDB(db, "stranger", "imagelocation4")
	updateRecordStatusToDB(db, id4, "reject")

	imageid, err := getLastedRecordImageID(db)
	fmt.Println(imageid)

	// records, err := getHisRecDBbyTime(db, "2021-03-25 12:00:00", "2021-03-25 13:50:00")
	// fmt.Println(records)
	// fmt.Println(err)

}
