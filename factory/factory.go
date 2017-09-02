package factory

// type Job interface {
// 	Do() error
// }

// var JobQueue chan Job

// type Worker struct {
// 	WorkerPool chan chan Job
// 	JobChannel chan Job

// 	quit chan struct{}
// }

// func NewWorker(workerPool chan chan Job) *Worker {
// 	return *Worker{
// 		WorkerPool: workerPool,
// 		JobChannel: make(chan Job),
// 		quit:       make(chan struct{}),
// 	}
// }

// func (w *Worker) Start() {
// 	go func() {
// 		for {
// 			// register the current worker into the worker queue.
// 			w.WorkerPool <- w.JobChannel

// 			select {
// 			case job := <-w.JobChannel:
// 				// we have received a work request.
// 				if err := job.Do(); err != nil {
// 					log.Errorf("Error uploading to S3: %s", err.Error())
// 				}

// 			case <-w.quit:
// 				// we have received a signal to stop
// 				return
// 			}
// 		}
// 	}()
// }

// func (w *Worker) Stop() {
// 	select {
// 	case <-w.quit:
// 	default:
// 		close(w.quit)
// 	}
// }

// type Dispatcher struct {
// 	// A pool of workers channels that are registered with the dispatcher
// 	WorkerPool chan chan Job
// }

// func NewDispatcher(maxWorkers int) *Dispatcher {
// 	pool := make(chan chan Job, maxWorkers)
// 	return &Dispatcher{WorkerPool: pool}
// }

// func (d *Dispatcher) Run() {
// 	// starting n number of workers
// 	for i := 0; i < d.maxWorkers; i++ {
// 		worker := NewWorker(d.pool)
// 		worker.Start()
// 	}

// 	go d.dispatch()
// }

// func (d *Dispatcher) dispatch() {
// 	for {
// 		select {
// 		case job := <-JobQueue:
// 			// a job request has been received
// 			go func(job Job) {
// 				// try to obtain a worker job channel that is available.
// 				// this will block until a worker is idle
// 				jobChannel := <-d.WorkerPool

// 				// dispatch the job to the worker job channel
// 				jobChannel <- job
// 			}(job)
// 		}
// 	}
// }
