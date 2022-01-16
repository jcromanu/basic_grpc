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
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatal("Error on client connection", err)
	}
	defer conn.Close()
	client := pb.NewRecordServiceClient(conn)
	//Simple RPC
	recordResponse, err := client.GetRecord(context.Background(), &pb.RecordRequest{Id: "1"})
	if err != nil {
		log.Fatal("Error on record", err)
	}
	log.Println("Simple RPC record :", recordResponse.Record.RecordId)
	//Server side streaming
	stream, err := client.ListRecords(context.Background(), &pb.User{UserId: "1", Type: pb.User_ENTERPRISE})
	if err != nil {
		log.Fatal("Error on streaming ", err)
	}
	for {
		recordResponse, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error on streaming ", err)
		}
		log.Print("server streaming Record id:", recordResponse.Record.RecordId)
	}
}
