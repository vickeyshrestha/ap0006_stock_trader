package stocktrader

import (
	nats "github.com/vickeyshrestha/sharing-services/drivers/nats"
	"log"
	"time"
)

type TimeSeries struct {
	ExchangeName string
	Stocks       Stock
	DateTime     time.Time
}

type Stock struct {
	Companies map[string]float64
}

var queue []TimeSeries

// Subscriber will subscribe to a specific topic in NATS and will process the enqueue feature
func Subscriber(natsUrl string) {
	encodedConnection, err := nats.NewNatsConnectionClient(natsUrl)
	if err != nil {
		log.Println(err)
		return
	}
	defer encodedConnection.Close()

	tsChanRecv := make(chan *TimeSeries)
	encodedConnection.BindRecvChan(SimNatsTopic, tsChanRecv)

	for {

		// get live data
		req := <-tsChanRecv
		log.Printf("Received request -> Exchange name: %s | Stocks: %v", req.ExchangeName, req.Stocks)

		// in reality, this timestamp should have come from the serverside which determines actual timestamp (future work in simulator).
		// time.Now() is not good since we should also count the delay that might happen between client and server
		// this can be later used to check the expired data
		req.DateTime = time.Now()

		// we will determine the size of the queue to be 15
		// to avoid hammering the memory, we have to limit the queue size
		if len(queue) < QueueSize {
			queue = enqueue(queue, *req)
		}
		log.Printf("current queue -> data: %v | size: %d", queue, len(queue))
	}
}

func Dequeue() {
	for range time.Tick(time.Second * DequeueProcessIntervalInSeconds) {
		if len(queue) > 0 {
			// delete data only when the first element data is older than 15 seconds
			dataDuration := time.Now().Sub(queue[0].DateTime)
			if dataDuration > time.Second*TimeToCalculateValuesAndDequeue {
				purgedData, newQueue := dequeue(queue)
				log.Printf("Dequeued | data: %v | alive duration: %v", purgedData, dataDuration)
				queue = newQueue
			}
			purgedData, newQueue := dequeue(queue)
			log.Println("Dequeue data: ", purgedData)
			queue = newQueue
		}
	}
}
