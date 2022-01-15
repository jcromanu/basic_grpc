package main

import (
	"context"
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
	recordResponse, err := client.GetRecord(context.Background(), &pb.RecordRequest{Id: "1"})
	if err != nil {
		log.Fatal("Error on record", err)
	}
	log.Println("Simple RPC record :", recordResponse.Record.RecordId)
}
