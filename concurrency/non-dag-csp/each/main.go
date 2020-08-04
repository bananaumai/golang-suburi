package main

import (
	"context"
	"fmt"
	"time"
)

const nWorkers int = 4
const nDownloaders int = 4

type Reference struct{ index int }
type Location struct{ index int }
type Content struct{ index int }
type LocContent struct {
	loc     Location
	content Content
}
type WorkerInput struct {
	loc      Location
	contents chan LocContent
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
	workerInputs chan<- WorkerInput,
) {
	for {
		ref := <-references
		loc := ref.resolveLocation()
		contents := make(chan LocContent, 1)
		input := WorkerInput{
			loc:      loc,
			contents: contents,
		}
		workerInputs <- input
		lc := <-contents
		processContent(ref, lc.content)
	}
}

func worker(inputs <-chan WorkerInput) {
	for input := range inputs {
		content := downloadContent(input.loc)
		input.contents <- LocContent{input.loc, content}
	}
}

func processReferences(references <-chan Reference) {
	workerInputs := make(chan WorkerInput)
	for i := 0; i < nWorkers; i++ {
		go worker(workerInputs)
	}
	for i := 0; i < nDownloaders; i++ {
		go downloader(references, workerInputs)
	}
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
