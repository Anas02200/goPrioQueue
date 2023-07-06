package queueimplem

import (
	"container/heap"
	jobs "main/queueimplem/jobmodels"
	"sort"
	"sync"
)

//first version doesent support concurrent opeations
// lets make it concurrent now

type AbsPrioQueue interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any
}

// change to struct later
type PriorityQueue[T jobs.BasicJob] struct {
	sync.Mutex
	Jobs []*T
}

func (pq PriorityQueue[T]) Len() int { return len(any(pq.Jobs).([]*jobs.BasicJob)) }
func (pq PriorityQueue[T]) Less(i, j int) bool {
	pq.Lock()
	defer pq.Unlock()
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	jobss := any(pq.Jobs).([]*jobs.BasicJob)

	return jobss[i].Priority > jobss[j].Priority
}
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq.Lock()
	defer pq.Unlock()
	jobss := any(pq.Jobs).([]*jobs.BasicJob)
	jobss[i], jobss[j] = jobss[j], jobss[i]
	jobss[i].Index = i
	jobss[j].Index = j
}
func (pq *PriorityQueue[T]) Push(x any) {
	pq.Lock()
	defer pq.Unlock()
	//######

	//nasty casting bc generics not mature enough hh
	jobss := any(pq.Jobs).([]*jobs.BasicJob)
	n := len(jobss)
	item := x.(*jobs.BasicJob)
	item.Index = n
	jobsll := append(jobss, item)
	pq.Jobs = any(jobsll).([]*T)

	//######
	//must call init to order the heap
	//heap.Push(pq, x)

}
func (pq *PriorityQueue[T]) Pop() any {
	pq.Lock()
	defer pq.Unlock()
	//###### buggy implem , call heap.pop directly in client code
	old := any(pq.Jobs).([]*jobs.BasicJob)
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety

	jobsll := old[0 : n-1]

	pq.Jobs = any(jobsll).([]*T)
	//#######

	//return heap.Pop(pq).(*jobs.BasicJob)
	return item
}
func (pq *PriorityQueue[T]) Update(item *jobs.BasicJob, value string, priority int32) {
	//update priority

	item.PromptData = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
