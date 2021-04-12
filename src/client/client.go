package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"github.com/CPEN391-Team-4/backend/src/environment"
	"google.golang.org/grpc"
)

const READ_BUF_SIZE = 16

const NUM_TEST_FRAMES = 4000
const CHUNK_NUM = 3
const CHUNK_SIZE = 20 * 1014

func verifyFace(client pb.RouteClient, ctx context.Context, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, READ_BUF_SIZE)

	var photo pb.Photo
	stream, err := client.VerifyUserFace(ctx)
	if err != nil {
		log.Fatalf("%v.VerifyUserFace(_) = _, %v", client, err)
	}
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
		req := pb.FaceVerificationReq{Photo: &photo}
		if err := stream.Send(&req); err != nil && err != io.EOF {
			log.Fatalf("%v.Send(%v) = %v", stream, &req, err)
		}
		sizeTotal += n
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v, %v", reply.Accept, reply.Confidence)

	return nil
}
func streamVideo(client pb.VideoRouteClient, ctx context.Context, file string, streaming chan bool) error {

	frame := pb.Frame{}
	stream, err := client.StreamVideo(ctx)
	if err != nil {
		log.Fatalf("%v.StreamVideo(_) = _, %v", client, err)
	}
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		loc := 0
		for {
			if loc + READ_BUF_SIZE > len(buf) {
				frame.Chunk = buf[loc:]
				frame.LastChunk = true
			} else {
				frame.Chunk = buf[loc:loc+READ_BUF_SIZE]
				frame.LastChunk = false
			}

			frame.Number = int32(i)

			req := pb.Video{Frame: &frame, Name: "Test"}
			if err := stream.Send(&req); err != nil && err != io.EOF {
				log.Fatalf("%v.Send(%v) = %v", stream, &req, err)
			}
			loc += READ_BUF_SIZE
			if loc > len(buf) {
				break
			}
		}
		if i == 0 && stream != nil {
			streaming <- true
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply)

	return nil
}

func getAllUserNames(c pb.RouteClient, ctx context.Context) error {
	users, err := c.GetAllUserNames(ctx, &pb.Empty{})
	if err != nil {
		return err
	}

	log.Printf("Route summary: %v", users)

	return nil
}

func reqToLock(c pb.RouteClient, ctx context.Context) error {
	stream, err := c.RequestToLock(ctx)
	if err != nil {
		return err
	}

	for {
		req, err := stream.Recv()
		log.Printf("req=%v", req)
		if err != nil {
			return err
		}
		stream.Send(&pb.LockConnection{Setup: true})
	}

	return nil
}
func reqToStream(c pb.VideoRouteClient, ctx context.Context) error {
	stream, err := c.RequestToStream(ctx)
	if err != nil {
		return err
	}

	for {
		req, err := stream.Recv()
		log.Printf("req=%v", req)
		if err != nil {
			return err
		}
		stream.Send(&pb.InitialConnection{Setup: true})
	}

	return nil
}

func addUser(client pb.RouteClient, ctx context.Context, file string, name string, restricted bool) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, READ_BUF_SIZE)

	var photo pb.Photo
	stream, err := client.AddTrustedUser(ctx)
	if err != nil {
		log.Fatalf("%v.AddTrustedUser(_) = _, %v", client, err)
	}
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
		req := pb.User{Photo: &photo, Name: name, Restricted: restricted}
		if err := stream.Send(&req); err != nil && err != io.EOF {
			log.Fatalf("%v.Send(%v) = %v", stream, &req, err)
		}
		sizeTotal += n
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply)
	return nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) []byte {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return []byte(string(b))
}

func streamVideoOld(client pb.VideoRouteClient, ctx context.Context) error {
	var frame pb.Frame

	stream, err := client.StreamVideo(ctx)
	if err != nil {
		log.Fatalf("%v.StreamVideo(_) = _, %v", client, err)
	}
	for i := 0; i < NUM_TEST_FRAMES; i++ {
		for j := 0; j < CHUNK_NUM; j++ {
			frame.Chunk = randSeq(CHUNK_SIZE)
			frame.LastChunk = j == CHUNK_NUM-1
			frame.Number = int32(i)
			req := pb.Video{Frame: &frame, Name: "Test"}
			if err := stream.Send(&req); err != nil && err != io.EOF {
				log.Fatalf("%v.Send(%v) = %v", stream, &req, err)
			}
			log.Printf("Sent frame.Number=%v, frame.LastChunk=%v", frame.Number, frame.LastChunk)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply)
	return nil
}

func sendPullVideo(client pb.VideoRouteClient, ctx context.Context) error {
	var frame pb.Frame
	sendStream, err := client.StreamVideo(ctx)
	if err != nil {
		log.Fatalf("%v.StreamVideo(_) = _, %v", client, err)
	}
	pullStream, err := client.PullVideoStream(ctx, &pb.PullVideoStreamReq{Id: "default"})
	if err != nil {
		log.Fatalf("%v.PullVideoStream(_) = _, %v", client, err)
	}

	for i := 0; i < NUM_TEST_FRAMES+1; i++ {
		if i < NUM_TEST_FRAMES+1 {
			for j := 0; j < 10; j++ {
				frame.Chunk = []byte{byte(j)}
				frame.LastChunk = j == 9
				frame.Number = int32(i)
				req := pb.Video{Frame: &frame, Name: "Test"}
				if err := sendStream.Send(&req); err != nil && err != io.EOF {
					log.Fatalf("%v.Send(%v) = %v", sendStream, &req, err)
				}

				log.Printf("Sent frame.Number=%v, frame.LastChunk=%v", frame.Number, frame.LastChunk)
			}
		}

		if i == 0 {
			continue
		}

		reply, err := pullStream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("%v.Recv() = %v", pullStream, err)
		}
		if int(reply.Video.Frame.Number) != i-1 {
			log.Fatalf("Wrong frame, expected %v, recieved %v", i, int(reply.Video.Frame.Number))
		}
		log.Printf("Recieved %v", int(reply.Video.Frame.Number))
		if reply.Closed {
			log.Printf("Stream closed")
			break
		}
	}
	_, err = sendStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", sendStream, err, nil)
	}

	return nil
}
func pullVideo(client pb.VideoRouteClient, ctx context.Context) error {
	pullStream, err := client.PullVideoStream(ctx, &pb.PullVideoStreamReq{Id: "default"})
	if err != nil {
		log.Fatalf("%v.PullVideoStream(_) = _, %v", client, err)
	}

	for {

		reply, err := pullStream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("%v.Recv() = %v", pullStream, err)
		}
		log.Printf("Recieved %v", reply)
		if reply.Closed {
			log.Printf("Stream closed")
			break
		}
	}
	err = pullStream.CloseSend()

	return err
}


func requestToStream(client pb.VideoRouteClient, ctx context.Context, file string) error {
	reqStream, err := client.RequestToStream(ctx)
	if err != nil {
		log.Fatalf("%v.RequestToStream(_) = _, %v", client, err)
	}
	up := make(chan bool)
	streaming := make(chan bool)

	go func() {
		for {
			streamNow := <- up
			if streamNow {
				log.Printf("streamVideo")
				err := streamVideo(client, ctx, file, streaming)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				log.Printf("Stream done")
				return
			}
		}
	}()

	go func() {
		for {
			reply, err := reqStream.Recv()

			up <- reply.Request

			<- streaming
			err = reqStream.Send(&pb.InitialConnection{Setup: true})
			log.Printf("Send setup true")
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	<-time.After(10 * time.Second)

	pullStream, err := client.PullVideoStream(ctx, &pb.PullVideoStreamReq{Id: "default"})
	if err != nil {
		log.Fatalf("%v.PullVideoStream(_) = _, %v", client, err)
	}

	for {
		reply, err := pullStream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("%v.Recv() = %v", pullStream, err)
		}
		if reply.Video != nil && reply.Video.Frame != nil {
			log.Printf("Recieved Frame.Number=%v", reply.Video.Frame.Number)
		}

		if reply.Closed {
			log.Printf("Stream closed")
			break
		}
	}

	return nil
}


func main() {
	environ := environment.Env{}
	environ.ReadEnv()

	conn, err := grpc.Dial(environ.ServerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRouteClient(conn)
	svc := pb.NewVideoRouteClient(conn)

	verifyFaceCmd := flag.NewFlagSet("verifyface", flag.ExitOnError)
	streamVideoCmd := flag.NewFlagSet("streamvideo", flag.ExitOnError)
	requestStreamVideoCmd := flag.NewFlagSet("requeststream", flag.ExitOnError)
	addUserCmd := flag.NewFlagSet("adduser", flag.ExitOnError)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if len(os.Args) < 2 {
		fmt.Println("expected subcommand 'verifyface' | 'addUser' | 'listusers' | 'streamvideo' | 'requeststream'")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "verifyface":
		verifyFaceCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'verifyface'")
		fmt.Println("  tail:", verifyFaceCmd.Args())
		if len(verifyFaceCmd.Args()) < 1 {
			fmt.Println("expected subcommand 'verifyface' FILE argument")
			os.Exit(1)
		}
		err = verifyFace(c, ctx, verifyFaceCmd.Args()[0])
	case "listusers":
		fmt.Println("subcommand 'listusers'")
		err = getAllUserNames(c, ctx)
	case "adduser":
		fmt.Println("subcommand 'addUser'")
		if len(addUserCmd.Args()) < 3 {
			fmt.Println("expected subcommand 'adduser' FILE, NAME, RESTRICTED argument")
			os.Exit(1)
		}
		restr := addUserCmd.Arg(2)
		resInt, err := strconv.Atoi(restr)
		if err != nil {
			os.Exit(1)
		}
		restricted := resInt != 0
		err = addUser(c, ctx, addUserCmd.Arg(0), addUserCmd.Arg(1), restricted)
		if err != nil {
			os.Exit(1)
		}
	case "streamvideo":
		fmt.Println("subcommand 'streamvideo'")
		_ = streamVideoCmd.Parse(os.Args[2:])
		fmt.Println("  tail:", streamVideoCmd.Args())
		if len(streamVideoCmd.Args()) < 1 {
			fmt.Println("expected subcommand 'streamvideo' FILE argument")
			os.Exit(1)
		}
		err = streamVideo(svc, ctx, streamVideoCmd.Args()[0], nil)
	case "requeststream":
		fmt.Println("subcommand 'requeststream'")
		_ = requestStreamVideoCmd.Parse(os.Args[2:])
		fmt.Println("  tail:", requestStreamVideoCmd.Args())
		if len(requestStreamVideoCmd.Args()) < 1 {
			fmt.Println("expected subcommand 'requeststream' FILE argument")
			os.Exit(1)
		}
		err = requestToStream(svc, ctx, requestStreamVideoCmd.Args()[0])
	case "reqtostream":
		fmt.Println("subcommand 'reqtostream'")
		err = reqToStream(svc, ctx)
	case "sendpullvideo":
		fmt.Println("subcommand 'sendpullvideo'")
		err = sendPullVideo(svc, ctx)
	case "pullvideo":
		fmt.Println("subcommand 'pullvideo'")
		err = pullVideo(svc, ctx)
	case "reqtolock":
		fmt.Println("subcommand 'reqtolock'")
		err = reqToLock(c, ctx)
	default:
		fmt.Println("expected subcommand")
		os.Exit(1)
	}

	if err != nil {
		os.Exit(1)
	}
}
