package main

import (
	"context"
	"fmt"

	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	_ "github.com/go-sql-driver/mysql"
)

const DeviceTable = "devices"

func (rs *routeServer) UpdateDeviceToken(ctx context.Context, device *pb.DeviceVerify) (*pb.Empty, error) {

	var err error

	find, err := rs.CheckDeviceExisted(device.Deviceid)
	if err != nil {
		return &pb.Empty{}, err
	}

	if find {
		fmt.Println("update old device token")
		err := rs.UpdateToken(device.Deviceid, device.Token)
		if err != nil {
			return &pb.Empty{}, err
		}
	} else {
		fmt.Println("add new device and its token")
		err := rs.AddTODeviceDB(device.Deviceid, device.Token)
		if err != nil {
			return &pb.Empty{}, err
		}

	}
	return &pb.Empty{}, err
}

//.
//.
//.
//.
//All the functions for the communication between server to database for history record part.

func (rs *routeServer) AddTODeviceDB(deviceID string, token string) error {

	sql_q := fmt.Sprintf(
		"INSERT INTO `%s` (deviceid, token ) VALUES ('%s', '%s');",
		DeviceTable, deviceID, token)
	_, err := rs.conn.Exec(sql_q)
	if err != nil {
		return err
	}

	return err
}

func (rs *routeServer) CheckDeviceExisted(deviceID string) (bool, error) {
	find := true
	sql_q := "SELECT exists( select * from " + DeviceTable + " WHERE deviceid = ?)"
	res, err := rs.conn.Query(sql_q, deviceID)

	if err != nil {
		return find, err
	}

	for res.Next() {
		var result int
		err = res.Scan(&result)
		if err != nil {
			panic(err.Error())
		}
		if result == 1 {
			find = true
		} else {
			find = false
		}
	}
	return find, err
}

func (rs *routeServer) UpdateToken(deviceID string, token string) error {
	sql_q := fmt.Sprintf(
		"UPDATE `%s` SET token = '%s' where deviceid = '%s';",
		DeviceTable, token, deviceID)
	_, err := rs.conn.Exec(sql_q)
	return err
}

func (rs *routeServer) GetAllTokens() ([]string, error) {
	sql_q := "SELECT token FROM " + DeviceTable
	res, err := rs.conn.Query(sql_q)

	if err != nil {
		return nil, err
	}

	var tokens []string

	for res.Next() {
		var token string
		err = res.Scan(&token)
		if err != nil {
			panic(err.Error())
		}
		tokens = append(tokens, token)
	}

	return tokens, err
}
