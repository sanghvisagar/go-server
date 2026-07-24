package main

import "time"

type request struct {
	timestamp time.Time

}

type DoubleEndedQueue struct {
	requests [request]
}

func addToQueue(req *request) {

}

func evictKeys(now time.Time, queue *DoubleEndedQueue) {
	for req := range(queue.requests) {
		if req.timestamp.isBefore(now - time.Second)
		delete(req)
	}
}

func deleteFromQueue() {

}



func main() {
	r1 := request{time.now}
	r2 := request{time.now - time.Second*2}
	r3 := request{time.now - time.Second*3}
	r4 := request{time.now - time.Second*4}
	r5 := request{time.now - time.Second}


	requestsInTimeFrame := DoubleEndedQueue{requests: {r1,r2,r3,r4,r5}}

	
	for req := range(requests) {
		addToQueue(&req)
		evictKeys(time.now, DoubleEndedQueue)

		if requests 
	}

}