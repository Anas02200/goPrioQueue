package main

import (
	"container/heap"
	"fmt"
	"log"
	"main/queueimplem"
	jobs "main/queueimplem/jobmodels"
	"main/queueserver"
	"time"
)

func main() {
	items := map[string]int{
		"banana": 30, "apple": 20, "pear": 40,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := &queueimplem.PriorityQueue[jobs.BasicJob]{}
	i := 0

	for k, v := range items {

		pq.Jobs = append(
			any(pq.Jobs).([]*jobs.BasicJob), &jobs.BasicJob{
				Id:         "anas",
				UserId:     "anas",
				PromptData: k,
				TimeIssued: time.Now(),
				Priority:   int32(v),
				Index:      i,
			},
		)

		i++

	}

	log.Println("ddd")
	//
	heap.Init(pq)
	//
	//// Insert a new item and then modify its priority.
	//
	//for true {
	//	//start server and for each request push to the prio queue
	//
	//}

	item := &jobs.BasicJob{
		Id:         "anas",
		UserId:     "anas",
		PromptData: "added value",
		TimeIssued: time.Now(),
		Priority:   60,
		Index:      50,
	}
	itemlow := &jobs.BasicJob{
		Id:         "anas",
		UserId:     "anas",
		PromptData: "added  v2",
		TimeIssued: time.Now(),
		Priority:   5000,
		Index:      i,
	}

	heap.Push(pq, item)
	heap.Push(pq, itemlow)
	//pq.Update(item, "new prompt", 22)
	//pq.Update(itemlow, "new prompt v2", 6000)
	//
	//// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*jobs.BasicJob)
		fmt.Printf("%.2d:%s\n ", item.Priority, item.PromptData)
	}

	queueserver.InitServer("5050")
}
