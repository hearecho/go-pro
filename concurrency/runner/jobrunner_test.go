package runner

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestJobRunner(t *testing.T) {
	fmt.Println("开始运行")
	r := NewJobRunner(time.Second * 5)
	r.Add(createTask(),createTask(),createTask())
	if err:= r.Start();err != nil {
		switch err {
		case ErrTimeout:
			log.Println("task stop cased by timout")
			os.Exit(1)
		case ErrInterrupt:
			log.Panicln("task stop cased by interrupt")
			os.Exit(2)
		}
	}
}

//模拟任务
func createTask() func(int) {
	return func(i int) {
		log.Printf("Processor -Task:%d",i)
		time.Sleep(time.Duration(i)*time.Second)
	}
}