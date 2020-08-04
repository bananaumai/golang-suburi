package main

import (
	"fmt"
	"sync"
	"time"
)

const nWorkers int = 4

type Reference struct{ index int }
type Location struct{ index int }
type Content struct{ index int }
type LocContent struct {
	loc     Location
	content Content
}

func (ref Reference) resolveLocation() Location {
	log("Resolving location for %#v", ref)
	return Location{ref.index}
}

func downloadContent(loc Location) Content {
	log("Downloading %#v", loc)
	time.Sleep(10 * time.Millisecond)
	return Content{loc.index}
}

func processContent(ref Reference, content Content) {
	log("Processing %#v %#v", ref, content)
}

func log(format string, a ...interface{}) {
	b := []interface{}{time.Now().Format(time.RFC3339)}
	fmt.Printf("%s "+format+"\n", append(b, a...)...)
}

func work(
	locations <-chan Location,
	contents chan<- LocContent,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	loc := <-locations
	content := downloadContent(loc)
	contents <- LocContent{loc, content}
}

func main() {
	locations := make(chan Location)
	contents := make(chan LocContent)

	mut := sync.Mutex{}
	requested := make(map[Location][]Reference)

	num := 1000

	wg := sync.WaitGroup{}
	wg.Add(num)
	go func() {
		for index := 0; index < num; index++ {
			mut.Lock()
			ref := Reference{index: index}
			loc := ref.resolveLocation()
			refs, present := requested[loc]
			if !present {
				requested[loc] = []Reference{ref}
				go work(locations, contents, &wg)
				locations <- loc
			} else {
				requested[loc] = append(refs, ref)
			}
			mut.Unlock()
		}
	}()

	go func() {
		for lc := range contents {
			mut.Lock()
			refs := requested[lc.loc]
			delete(requested, lc.loc)
			mut.Unlock()
			for _, ref := range refs {
				processContent(ref, lc.content)
			}
		}
	}()

	wg.Wait()

	//time.Sleep(1 * time.Second)
}
