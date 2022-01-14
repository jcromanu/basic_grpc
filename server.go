package main

import (
	context "context"

	pb "example.com/record"
)

type RecordServiceServer struct {
}

func (r *RecordServiceServer) GetRecord(ctx context.Context, rq *pb.RecordRequest) (*pb.RecordResponse, error) {
	for record := range recordList {
			if rq.GetId() == record. {
			}
	}
}

//Server side streaming RPC
func (r *RecordServiceServer) ListRecords(u *pb.User, rs pb.RecordService_ListRecordsServer) error {

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
func (r *RecordServiceServer) mustEmbedUnimplementedRecordServiceServer() {
}

var recordList = []*pb.Record{{RecordId: "1", UserId: "1", Volume: 1}}
