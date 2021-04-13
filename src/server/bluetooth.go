package main

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	_ "github.com/go-sql-driver/mysql"
)

const BlueToothTable = "bluetooth"

//update the de1 id and username
func (rs *routeServer) SendDe1ID(ctx context.Context, in *pb.BluetoothInfo) (*pb.Empty, error) {
	fmt.Println("receive de1 id and username.")

	err := rs.addBlueinfoTODB(in.Username, in.De1ID)
	if err != nil {
		fmt.Println(err)
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

//get the de1 id and username
func (rs *routeServer) GetDe1ID(ctx context.Context, in *pb.MainUser) (*pb.BluetoothInfo, error) {
	fmt.Println("receive request to get de1 id.")
	id, err := rs.getDe1IDFromDB(in.Username)

	info := &pb.BluetoothInfo{De1ID: id}
	if err != nil {
		fmt.Println(err)
		return &pb.BluetoothInfo{}, err
	}
	return info, nil
}

//add the username and the de1device id into the bluetooth table
func (rs *routeServer) addBlueinfoTODB(username string, de1id string) error {
	sql_q := fmt.Sprintf(
		"INSERT INTO `%s` (username, de1id ) VALUES ('%s', '%s');",
		BlueToothTable, username, de1id)
	_, err := rs.conn.Exec(sql_q)
	if err != nil {
		return err
	}

	return err
}

//by giving the username, get the device id of that username.
func (rs *routeServer) getDe1IDFromDB(username string) (string, error) {
	sql_q := "SELECT de1id FROM " + BlueToothTable + " where username = " + "'" + username + "'"
	fmt.Println(sql_q)
	res, err := rs.conn.Query(sql_q)

	if err != nil {
		return "", err
	}

	var de1id string

	for res.Next() {
		err = res.Scan(&de1id)
		if err != nil {
			panic(err.Error())
		}
	}

	return de1id, err
}
