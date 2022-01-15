package main

import (
	context "context"
	"log"
	"net"

	pb "example.com/record"
	"google.golang.org/grpc"
)

type RecordServiceServer struct {
	pb.UnimplementedRecordServiceServer
}

var recordList = []pb.Record{{RecordId: "1", UserId: "1", Volume: 1}}

func (r *RecordServiceServer) GetRecord(ctx context.Context, rq *pb.RecordRequest) (*pb.RecordResponse, error) {
	for _, record := range recordList {
		if rq.GetId() == record.RecordId {
			return &pb.RecordResponse{Record: &record, Error: nil}, nil
		}
	}
	return nil, nil
}

//Server side streaming RPC
func (r *RecordServiceServer) ListRecords(u *pb.User, rs pb.RecordService_ListRecordsServer) error {
	return nil
}

//Client side streaming RPC
func (r *RecordServiceServer) SetRecords(rs pb.RecordService_SetRecordsServer) error {
	return nil
}

//Bi directional side streaming RPC
func (r *RecordServiceServer) RecordPong(rs pb.RecordService_RecordPongServer) error {
	return nil
}

//TBD ???
func (r *RecordServiceServer) mustEmbedUnimplementedRecordServiceServer() {}

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("failes to listen", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRecordServiceServer(grpcServer, &RecordServiceServer{})
	grpcServer.Serve(lis)
}
