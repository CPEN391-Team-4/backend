package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	_ "github.com/go-sql-driver/mysql"
)

// Table for devices
const DeviceTable = "devices"

/********************* DATABASE *********************/

// AddTODeviceDB Add new entry to device table
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

// CheckDeviceExisted Check if the device id is already existed in the device table
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

// UpdateToken Update the device token of the a certain device id
func (rs *routeServer) UpdateToken(deviceID string, token string) error {
	sql_q := fmt.Sprintf(
		"UPDATE `%s` SET token = '%s' where deviceid = '%s';",
		DeviceTable, token, deviceID)
	_, err := rs.conn.Exec(sql_q)
	return err
}

// GetAllTokens Gget all the tokens in the device table
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

/********************* DEVICE API *********************/

// UpdateDeviceToken update the device token of the a certain device
func (rs *routeServer) UpdateDeviceToken(ctx context.Context, device *pb.DeviceVerify) (*pb.Empty, error) {

	var err error

	find, err := rs.CheckDeviceExisted(device.Deviceid)
	if err != nil {
		return &pb.Empty{}, err
	}

	if find {
		// Token already existed, update the token,
		log.Println("update old device token")
		err := rs.UpdateToken(device.Deviceid, device.Token)
		if err != nil {
			return &pb.Empty{}, err
		}
	} else {
		// Add a new entry into the device table
		log.Println("add new device and its token")
		err := rs.AddTODeviceDB(device.Deviceid, device.Token)
		if err != nil {
			return &pb.Empty{}, err
		}

	}
	return &pb.Empty{}, err
}