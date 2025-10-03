package main

import (
	"k8s.io/client-go/util/workqueue"
)

func main() {
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	// shutdown when process ends
	defer queue.ShutDown()
}
