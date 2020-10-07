package resourcePool

import (
	"errors"
	"io"
	"log"
	"sync"
)

/**
共享资源池
1. 因为并发共享资源，所以必须使用锁进行控制.
2. 资源需要关闭，所以资源最好实现Closer接口
3. 池概念肯定需要一个最大资源数量
4. 池肯定还需要向里面添加资源.
根据以上原则设计一个资源池
*/
type Pool struct {
	//并发安全
	m        sync.Mutex
	//资源,使用缓冲通道
	resources chan io.Closer
	//添加资源的函数
	factory  func() (io.Closer, error)
	//资源池的状态
	closed   bool
}

var ErrPoolClosed = errors.New("resource pool has closed")

func New(fn func()(io.Closer, error),size uint)(*Pool,error)  {
	if size < 0 {
		return nil,errors.New("size need more than 0")
	}
	return &Pool{
		resources: make(chan io.Closer,size),
		factory:  fn,
	},nil
}
/**
从资源池中获取资源
 */
func (p *Pool)Acquire()(io.Closer, error)  {
	select {
	case r,ok:= <-p.resources:
		log.Println("acquire:","shared resource")
		if !ok {
			return nil,ErrPoolClosed
		}
		return r,nil
	default:
		log.Println("acquire:","new resource")
		return p.factory()
	}
}
/**
释放资源
 */
func (p *Pool)Release(r io.Closer)  {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		_ = r.Close()
		return
	}
	select {
	case p.resources <- r:
		log.Println("Release:","In Queue")
	default:
		log.Println("Release:","Closing")
		_ = r.Close()
	}
}

func (p *Pool)Close()  {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resources)
	for r := range p.resources {
		_ = r.Close()
	}
}