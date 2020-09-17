package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"os/exec"
	"time"
)

func main() {
	flag.Parse()
	command := flag.Arg(0)
	log.Printf("command: %s\n", command)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//ctx := context.Background()

	cmd := exec.CommandContext(ctx, command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("failed to get stdout: %s", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("failed to get stderr: %s", err)
	}
	r := io.MultiReader(stderr, stdout)
	scanner := bufio.NewScanner(r)
	buf := make([]byte, 256)
	scanner.Buffer(buf, 1)

	go func() {
		log.Printf("start to scan")
		for scanner.Scan() {
			log.Print(scanner.Text())
		}
		log.Printf("finished to scan")
	}()

	//if err := cmd.Start(); err != nil {
	//	log.Fatalf("failed to start command: %s", err)
	//}
	//log.Printf("wait for cmd to be finished")
	//if err := cmd.Wait(); err != nil {
	//	log.Printf("wait failed: %s", err)
	//} else {
	//	log.Print("completed")
	//}
	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to run: %s", err)
	}
}
