package main

import (
	"context"
	"fmt"
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

func downloader(
	references <-chan Reference,
	locations chan<- Location,
	contents <-chan LocContent,
) {
	requested := make(map[Location][]Reference)
	for {
		select {
		case lc := <-contents:
			refs := requested[lc.loc]
			delete(requested, lc.loc)
			for _, ref := range refs {
				processContent(ref, lc.content)
			}
		default:
		}

		select {
		case ref := <-references:
			loc := ref.resolveLocation()
			refs, present := requested[loc]
			if !present {
				requested[loc] = []Reference{ref}
				locations <- loc
			} else {
				requested[loc] = append(refs, ref)
			}
		}
		//log("loop")
	}
}

func worker(
	locations <-chan Location,
	contents chan<- LocContent,
) {
	for loc := range locations {
		content := downloadContent(loc)
		contents <- LocContent{loc, content}
	}
}

func processReferences(references <-chan Reference) {
	locations := make(chan Location)
	contents := make(chan LocContent, 1)
	for i := 0; i < nWorkers; i++ {
		go worker(locations, contents)
	}
	go downloader(references, locations, contents)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	references := make(chan Reference)
	processReferences(references)
	for index := 1; ; index++ {
		select {
		case <-ctx.Done():
			log("timed out")
			return
		case references <- Reference{index}:
		}
	}
}
