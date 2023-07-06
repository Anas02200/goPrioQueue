package queueimplem

import (
	"container/heap"
	"fmt"
)

// Mostly copied from basic heap priority queue example

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue2 implements heap.Interface and holds Items.
type PriorityQueue2 []*Item

func (pq PriorityQueue2) Len() int { return len(pq) }

func (pq PriorityQueue2) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue2) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue2) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func queue(in <-chan *Item, out chan<- *Item) {
	// Make us a queue!
	pq := make(PriorityQueue2, 0)
	heap.Init(&pq)

	var currentItem *Item       // Our item "in hand"
	var currentIn = in          // Current input channel (may be nil sometimes)
	var currentOut chan<- *Item // Current output channel (starts nil until we have something)

	defer close(out)

	for {
		select {
		// Read from the input
		case item, ok := <-currentIn:
			if !ok {
				// The input has been closed. Don't keep trying to read it
				currentIn = nil
				// If there's nothing pending to write, we're done
				if currentItem == nil {
					return
				}
				continue
			}

			// Were we holding something to write? Put it back.
			if currentItem != nil {
				heap.Push(&pq, currentItem)
			}

			// Put our new thing on the queue
			heap.Push(&pq, item)

			// Turn on the output queue if it's not turned on
			currentOut = out

			// Grab our best item. We know there's at least one. We just put it there.
			currentItem = heap.Pop(&pq).(*Item)

			// Write to the output
		case currentOut <- currentItem:
			// OK, we wrote. Is there anything else?
			if len(pq) > 0 {
				// Hold onto it for next time
				currentItem = heap.Pop(&pq).(*Item)
			} else {
				// Oh well, nothing to write. Is the input stream done?
				if currentIn == nil {
					// Then we're done
					return
				}

				// Otherwise, turn off the output stream for now.
				currentItem = nil
				currentOut = nil
			}
		}
	}
}

func Main() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	in := make(chan *Item, 10) // Big input buffer and unbuffered output should give best sort ordering.
	out := make(chan *Item)    // But the system will "work" for any particular values

	// Start the queuing engine!
	go queue(in, out)

	// Stick some stuff on in another goroutine
	go func() {
		i := 0
		for value, priority := range items {
			in <- &Item{
				value:    value,
				priority: priority,
				index:    i,
			}
			i++
		}
		close(in)
	}()

	// Read the results
	for item := range out {
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
	fmt.Println()
}
