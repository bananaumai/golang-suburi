package main

import (
	"log"
	"sync"
	"time"
)

type (
	stateful struct {
		counter int
		mutex   sync.Mutex
	}

	container struct {
		ss    map[string]*stateful
		mutex sync.Mutex
	}
)

func (s *stateful) inc() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.counter++
}

func (s *stateful) dec() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.counter--
}

func (s *stateful) count() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.counter
}

func newContainer() *container {
	return &container{
		ss: make(map[string]*stateful),
	}
}

func (c *container) getOrAdd(name string) *stateful {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	s, ok := c.ss[name]
	if !ok {
		s = &stateful{}
		c.ss[name] = s
		log.Printf("created a stateful for %s", name)
	} else {
		log.Printf("found a stateful for %s", name)
	}
	s.inc()
	return s
}

func (c *container) delete(name string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, ok := c.ss[name]; ok {
		delete(c.ss, name)
	} else {
		panic("no stateful found for " + name)
	}
}

func main() {
	c := newContainer()
	names := []string{"foo", "foo", "bar"}

	log.Printf("test with %+v", names)
	var wg sync.WaitGroup
	for i, name := range names {
		i := i
		name := name
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			log.Printf("run %s(%d)", name, i)

			s := c.getOrAdd(name)

			s.dec()
			time.Sleep(1 * time.Second) // emulate some heavy task
			if s.count() <= 0 {
				log.Printf("try to delete: %s(%d)", name, i)
				c.delete(name)
			}
		}(&wg)
	}
	wg.Wait()
}
