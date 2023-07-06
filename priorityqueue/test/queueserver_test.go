package test

import (
	"context"
	prot "generated_proto"
	"main/queueimplem"
	"main/queueserver"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				queueimplem.Main()
			},
		)
	}
}

//todo for later

//func Test_queueserver_AddJobToQueue(t *testing.T) {
//	type fields struct {
//		UnimplementedQueueOpsServer prot.UnimplementedQueueOpsServer
//	}
//	type args struct {
//		in0 context.Context
//		in1 *prot.AddJobRequest
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *prot.AddJobReply
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{
//			name: "test1",
//			fields: fields{
//
//				UnimplementedQueueOpsServer: prot.UnimplementedQueueOpsServer{},
//			},
//			args: args{
//				context.Background(),
//				&prot.AddJobRequest{
//					Priority: 0,
//					Message:  "hello",
//				},
//			},
//			want: &prot.AddJobReply{
//				JobId: "lmao",
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(
//			tt.name, func(t *testing.T) {
//				qu := queueserver.QueueServer{
//					UnimplementedQueueOpsServer: tt.fields.UnimplementedQueueOpsServer,
//				}
//				got, err := qu.AddJobToQueue(tt.args.in0, tt.args.in1)
//
//				if (err != nil) != tt.wantErr {
//					t.Errorf("AddJobToQueue() error = %v, wantErr %v", err, tt.wantErr)
//					return
//				}
//				if !reflect.DeepEqual(got, tt.want) {
//					t.Errorf("AddJobToQueue() got = %v, want %v", got, tt.want)
//				}
//			},
//		)
//	}
//}

func Test_queueserver_GetJobsFromQueue(t *testing.T) {
	type fields struct {
		UnimplementedQueueOpsServer prot.UnimplementedQueueOpsServer
	}
	type args struct {
		in0 context.Context
		in1 *prot.GetNJobs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *prot.GetNJobsReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				qu := queueserver.QueueServer{
					UnimplementedQueueOpsServer: tt.fields.UnimplementedQueueOpsServer,
				}
				got, err := qu.GetJobsFromQueue(tt.args.in0, tt.args.in1)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetJobsFromQueue() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetJobsFromQueue() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
