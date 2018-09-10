package work_queue

//var wg sync.WaitGroup

type Worker interface {
	Run() interface{ }

}

type WorkQueue struct {
	Jobs         chan Worker
	Results      chan interface{}
	StopRequests chan int
	NumWorkers   uint
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	q.Jobs = make(chan Worker, maxJobs)
	q.NumWorkers = nWorkers
	q.StopRequests = make(chan int, int(nWorkers))
	q.Results = make(chan interface{}, maxJobs)
	//wg.Add(int(maxJobs))
	for i := uint(0); i < nWorkers; i++ {
		go q.worker()

	}
	//wg.Wait()
	//close(q.Jobs)

	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	numWorkers := queue.NumWorkers
	running := true
	// Run tasks from the queue, unless we have been asked to stop.
	for running {
		// TODO: run tasks from Jobs
		jobs := queue.Jobs
		stopRequests := queue.StopRequests
		results := queue.Results
		if len(stopRequests) > 0 {
			<-stopRequests
			running = false
		} else if len(jobs) > 0 {
			j := <-jobs
			if len(stopRequests) < int(numWorkers) {
				results <- j.Run()
			}
		} else {
			running = false
		}

	}
	//wg.Done()
}

func (queue WorkQueue) Enqueue(work Worker) {
	// TODO
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	// TODO
	for i := uint(0); i < queue.NumWorkers; i++ {
		queue.StopRequests <- 1
	}
}
