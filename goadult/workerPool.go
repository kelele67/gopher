// http://quant67.com/post/go/go-thread-pool.html
// golong 线程池的实现
// 并不是真正的线程，而是协程
// 假设要接受一个 POST 请求，并将请求中的 payload 数组部分依次进行处理
// 本程序将 payload 依次取出，包装后发送到一个叫 JobQueue 的 chan 中，线程池会处理提交的 Job

func payLoadHandler(w http.ResponseWriter, r *http.Request) {
	if s.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// read the body into a string for json decoding
	var content = &PayloadCollection{}
	err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// go through each payload and queue items individually to be posted to Worker
	fot _, payload := range content.Payloads {
		//let's create a job with the payload
		work := Job{Payload: payload}

		// push the work onto the queue
		JobQueue <- work
	}

	w.WriteHeader(http.StatusOK)
}

// 线程池实现
var (
		MaxWorker = os.Getenv("MAX_WORKERS")
		MaxQueue = ps.Getenv("MAX_QUEUE")
		)

// job represents the job to be run
type Job struct {
	Payload Payload
}

// a buffered channel that we can send work requests on
var JobQueue chan job

// worker represents the worker that executes the job
type Worker struct {
	// WorkerPool 是一个指向全局唯一的 chan 的引用，
	// 负责传递 worker 接收 job 的 chan
	// worker 空闲时，将自己的 jobchannel 中取出一个 chan ，并将 job 放入其中
	// 此时 worker 将从 chan 中接收到 job 并进行处理
	WorkerPool chan chan Job

	// woker 用于接收 job 的 chan
	JobChannel chan job

	// 用于给 worker 发送控制命令的 chan，用于停止 chan
	quit chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker {
		WorkerPool: workerPool
		JobChannel: make(chan Job)
		quit: make(chan bool)
	}
}

// start method starts the run loop for the worker, listening for quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func () {
		for {
			// register the current worker into the worker queue
			w.WorkerPool <- w.JobChannel

			select {
			case job := <- w.JobChannel:
				// we have received a work request
				// do the job
			case <- w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// stop signals the worker to stop listening for work requests
func (w Worker) Stop() {
	go func () {
		w.quit <- true
	}()
}

// alloc job
type Dispatcher struct {
	// a pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	// starting n number of eorkers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.pool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			go func (job Job) {
				// try to obtain a worker job channel that is available
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}

dispatcher := NewDispatcher(MaxWorker)
dispatcher.Run()