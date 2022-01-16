package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "example.com/record"
)

func main() {

	var opts []grpc.DialOption
	var recordRequestList = []*pb.RecordRequest{{Id: "1"}, {Id: "2"}}
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatal("Error on simple rpc ", err)
	}
	defer conn.Close()
	client := pb.NewRecordServiceClient(conn)
	//Simple RPC
	recordResponse, err := client.GetRecord(context.Background(), &pb.RecordRequest{Id: "1"})
	if err != nil {
		log.Fatal("Error on simple rpc ", err)
	}
	log.Println("Simple RPC record :", recordResponse.Record.RecordId)
	//Server side streaming
	stream, err := client.ListRecords(context.Background(), &pb.User{UserId: "1", Type: pb.User_ENTERPRISE})
	if err != nil {
		log.Fatal("Error on server streaming ", err)
	}
	for {
		recordResponse, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error on  server streaming ", err)
		}
		log.Print("server streaming Record id:", recordResponse.Record.RecordId)
	}
	//Client side streaming
	str, err := client.SetRecords(context.Background())
	if err != nil {
		log.Fatal("Error on client side streaming ", err)
	}
	for _, rr := range recordRequestList {
		if err := str.Send(rr); err != nil {
			log.Fatal("Error on client streaming", err)
		}
	}
	reply, err := str.CloseAndRecv()
	if err != nil {
		log.Fatal("Error on client streaming ", err)
	}
	log.Println("client side streaming: ", reply.Message)
}
