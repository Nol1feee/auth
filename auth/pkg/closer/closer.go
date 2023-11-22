package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

func NewCloser() *Closer {
	return &Closer{}
}

type Closer struct {
	Mu    sync.Mutex
	Once  sync.Once
	done  chan struct{}
	funcs []func() error
}

func (c *Closer) New(sig ...os.Signal) {
	c.done = make(chan struct{})

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, sig...)
	<-ch
	signal.Stop(ch)
	c.closeAll()
}

func (c *Closer) Add(f ...func() error) {
	c.Mu.Lock()
	c.funcs = append(c.funcs, f...)
	c.Mu.Unlock()
}

func (c *Closer) Wait() {
	<-c.done
}

// CloseAll calls all closer functions
func (c *Closer) closeAll() {
	c.Once.Do(func() {
		defer close(c.done)

		c.Mu.Lock()
		funcs := c.funcs
		c.funcs = nil
		c.Mu.Unlock()

		// call all Closer funcs async
		errs := make(chan error, len(funcs))
		for _, f := range funcs {
			go func(f func() error) {
				errs <- f()
			}(f)
		}

		for i := 0; i < cap(errs); i++ {
			if err := <-errs; err != nil {
				log.Println("error returned from Closer")
				break
			}
		}
	})
}

//var globalCloser = New()

//func Add(f ...func() error) {
//	globalCloser.Add(f...)
//}

//func Wait() {
//	globalCloser.Wait()
//}

//func CloseAll() {
//	globalCloser.CloseAll()
//}
