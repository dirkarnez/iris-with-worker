package main

import (
	"k8s.io/client-go/util/workqueue"
)

func main() {
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	// shutdown when process ends
	defer queue.ShutDown()

	queue.Add(key)
	queue.Forget(key)
	queue.AddRateLimited(key)
	queue.NumRequeues(key)
	key, quit := queue.Get()
	if quit {
		return false
	}
	defer queue.Done(key)
}
