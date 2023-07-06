package queueserver

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"main/queueimplem"
	jobs "main/queueimplem/jobmodels"
	"main/queueservices"
	"net"
	"sync"
	"time"

	prot "generated_proto"
)

type QueueServer struct {
	prot.UnimplementedQueueOpsServer
	qservice *queueservices.QService
}

func (qsrvr QueueServer) AddJobToQueue(ctx context.Context, req *prot.AddJobRequest) (
	*prot.AddJobReply, error,
) {

	// inject a service here do smth with the request generate id put in db then return smth // the service has a queue in it

	//reply := &prioqueue.AddJobReply{JobId: "ppp"}

	queue, err := qsrvr.qservice.AddToQueue(req)

	return queue, err

}
func (qsrvr QueueServer) GetJobsFromQueue(ctx context.Context, req *prot.GetNJobs) (
	*prot.GetNJobsReply, error,
) {

	// service with a queue implem here => pop from queue and do smth with em (for now )
	//( post to stable dif api ? client will be polling from queue

	queue, err := qsrvr.qservice.GetFromQueue(req)
	return queue, err
}

func InitServer(port string) {

	// maybe 2 servers one for get and one for add  and 2 separate rpc services

	listen, err := net.Listen("tcp", ":"+port)

	if err != nil {
		panic(err)
	}

	serv := grpc.NewServer()

	initjobs := []*jobs.BasicJob{
		{
			Id:         "anas",
			UserId:     "anas",
			PromptData: "added value",
			TimeIssued: time.Now(),
			Priority:   60,
			Index:      50,
		},
	}

	queue := &queueimplem.PriorityQueue[jobs.BasicJob]{
		Mutex: sync.Mutex{},
		Jobs:  initjobs,
	}

	srv := &QueueServer{

		qservice: queueservices.InitQservice(queue),
	}
	prot.RegisterQueueOpsServer(
		serv, srv,
	)

	//fromQueue, _ := srv.qservice.GetFromQueue(&prot.GetNJobs{NumOfJobs: 10})
	//
	//for i := 0; i < len(fromQueue.Jobs); i++ {
	//
	//	log.Println("######## INIT QUEUE WITH DATA ##############")
	//
	//	log.Println(fromQueue.Jobs[i].Priority)
	//	log.Println(fromQueue.Jobs[i].JobId)
	//	log.Println(fromQueue.Jobs[i].PromptData)
	//	log.Println(fromQueue.Jobs[i].JobId)
	//
	//	log.Println("######## END OF INITIAL QUEUE  DATA ##############")
	//}

	log.Println(serv.GetServiceInfo())

	if err := serv.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
