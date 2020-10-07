package resourcePool

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	maxGoroutines = 5
	poolResources = 2
)

type dbConnection struct {
	ID int32
}

var idCounter int32

func (db *dbConnection) Close() error {
	log.Println("Close: Connection", db.ID)
	return nil
}

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

func Test_Pool(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)
	p, err := New(createConnection, poolResources)
	if err != nil {
		log.Println(err)
	}
	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q,p)
		}(query)
	}
	wg.Wait()
	log.Println("shut down program")
	p.Close()
}

func performQueries(q int, p *Pool) {
	conn,err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)
	time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n",q,conn.(*dbConnection).ID)
}

