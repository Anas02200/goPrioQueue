package queueservices

import (
	"container/heap"
	prot "generated_proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"main/queueimplem"
	jobs "main/queueimplem/jobmodels"
	"math/rand"
	"sync"
	"time"
)

var lock sync.Mutex

type QService struct {
	queue queueimplem.AbsPrioQueue
}

// init method

// receiver methods with go_proto types

func InitQservice(queue queueimplem.AbsPrioQueue) (qservice *QService) {

	return &QService{queue: queue}

}

func (qs *QService) AddToQueue(in *prot.AddJobRequest) (*prot.AddJobReply, error) {

	//create BasicJob from addjob request
	// push it to the queue

	job := &jobs.BasicJob{
		Id:         uuid.NewString(),
		UserId:     "Anas",
		PromptData: in.Message,
		TimeIssued: time.Now(),
		Priority:   in.Priority,
		Index:      rand.Int(),
	}
	lock.Lock()

	heap.Push(qs.queue, job)
	//qs.queue.Push(job) this bad and buggy

	lock.Unlock()
	// do smth
	return &prot.AddJobReply{JobId: job.Id}, nil
}

func (qs *QService) GetFromQueue(in *prot.GetNJobs) (*prot.GetNJobsReply, error) {
	// do smth
	if qs.queue.Len() != 0 {

		//pop := qs.queue.Pop() very bad and buggy
		lock.Lock()
		pop := heap.Pop(qs.queue)
		lock.Unlock()
		basicJobspoped := pop.(*jobs.BasicJob)

		priority := basicJobspoped.Priority
		id := basicJobspoped.Id
		userId := basicJobspoped.UserId
		data := basicJobspoped.PromptData

		p := &prot.SystemJob{

			JobId:      id,
			UserId:     userId,
			PromptData: data,
			Priority:   priority,
		}

		return &prot.GetNJobsReply{
			Jobs: []*prot.SystemJob{
				p,
			},
		}, nil

	} else {

		return nil, status.Errorf(codes.NotFound, "Empty queue")
	}

}
