package main

import (
	context "context"
	"io"
	"log"
	"net"

	pb "example.com/record"
	"google.golang.org/grpc"
)

type RecordServiceServer struct {
	pb.UnimplementedRecordServiceServer
}

var recordList = []pb.Record{{RecordId: "1", UserId: "1", Volume: 1}, {RecordId: "2", UserId: "1", Volume: 1}}

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
	for _, record := range recordList {
		if err := rs.Send(&pb.RecordResponse{Record: &record, Error: nil}); err != nil {
			return err
		}
	}
	return nil
}

//Client side streaming RPC
func (r *RecordServiceServer) SetRecords(rs pb.RecordService_SetRecordsServer) error {
	var records string
	for {
		recordR, err := rs.Recv()
		if err == io.EOF {

			return rs.SendAndClose(&pb.Error{Code: 200, Message: records})
		}
		if err != nil {
			return err
		}
		records += " record: " + recordR.GetId()
	}
}

//Bi directional side streaming RPC
func (r *RecordServiceServer) RecordPong(rs pb.RecordService_RecordPongServer) error {
	for {
		in, err := rs.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err := rs.Send(&pb.RecordResponse{Record: &pb.Record{RecordId: in.Id, UserId: "1", Volume: 2}, Error: nil}); err != nil {
			return err
		}
	}
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
