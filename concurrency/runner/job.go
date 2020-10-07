package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

/**
使用timeout进行自监督，防止单一任务运行时间过长，导致阻塞，因为task是在循环中运行，是一个一个运行的。
不使用协程运行的原因是为了可控，控制每一个协程的运行时间，也可以知道哪些任务出现了超时的问题。
 */
type JobRunner struct {
	interrupt chan os.Signal
	complete chan error
	timeout <-chan time.Time
	tasks []func(int)
}

var ErrTimeout  = errors.New("received timeout")
var ErrInterrupt  = errors.New("received interrupt")
func NewJobRunner(t time.Duration) *JobRunner  {
	return &JobRunner{
		interrupt: make(chan os.Signal,1),
		complete:  make(chan error),
		timeout:   time.After(t),
	}
}

func (r *JobRunner)Add(tasks ...func(int))  {
	r.tasks = append(r.tasks,tasks...)
}

func (r *JobRunner)Start() error  {
	//接受系统中所有的信号
	signal.Notify(r.interrupt,os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <- r.complete :
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *JobRunner) run() error {
	for id,task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}
//查看是否是被中断
func (r *JobRunner) gotInterrupt() bool {
	select {
	case <- r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
