package main

import (
	"context"
	"fmt"
	prot "generated_proto"

	"log"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	//"context"
	//"fmt"
	//swagger "generated-client"
)

var wgroup sync.WaitGroup

func main() {

	conn, _ := grpc.Dial("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := prot.NewQueueOpsClient(conn)

	// supports concurrent writes for now
	for i := 0; i < 10; i++ {
		wgroup.Add(1)
		go addToQ(client, &wgroup, i)
	}
	wgroup.Wait()
	for i := 0; i < 10; i++ {
		getFromQ(client)

	}

}

func getFromQ(client prot.QueueOpsClient) {

	jobsReply, err := client.GetJobsFromQueue(context.Background(), &prot.GetNJobs{NumOfJobs: 10})
	if err != nil {
		log.Fatalf("failed to fetch: %v \n", err)
	} else {
		for i := 0; i < len(jobsReply.Jobs); i++ {
			log.Println("######## RECEIVED QUEUE  DATA ##############")

			log.Println(jobsReply.Jobs[i].Priority)
			log.Println(jobsReply.Jobs[i].JobId)
			log.Println(jobsReply.Jobs[i].PromptData)
			log.Println(jobsReply.Jobs[i].UserId)

			log.Println("######## END OF RECEIVED QUEUE  DATA ##############")

		}
	}
}

func addToQ(client prot.QueueOpsClient, group *sync.WaitGroup, i int) {
	defer group.Done()
	generatedId, err := client.AddJobToQueue(
		context.Background(), &prot.AddJobRequest{
			Priority: rand.Int31n(500),
			Message:  fmt.Sprintf("%s%d \n", "hello job num : ", i),
		},
	)
	if err != nil {
		log.Fatalf("failed to add to queue: %v \n", err)
	} else {

		log.Printf("generated id %s \n", generatedId.JobId)
	}
}
