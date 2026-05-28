package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

type FileJob struct {
	fileName string
}

type Result struct {
	fileName string
	words    int
}

// Producer

func producer(files []string) <-chan FileJob {
	out := make(chan FileJob)

	go func() {
		defer close(out)
		for _, val := range files {
			time.Sleep(time.Second * 1)
			fmt.Println("Doing some fake processing")
			out <- FileJob{fileName: val}
		}
	}()
	return out
}

// Consumer

func consumer(id int, ctx context.Context, wg *sync.WaitGroup, jobs <-chan FileJob, results chan Result) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-jobs:
			fmt.Printf("Worker %d processing ", id)
			if !ok {
				return
			}
			time.Sleep(1 * time.Millisecond)
			wordLen := len(strings.Split(val.fileName, "_"))
			results <- Result{
				fileName: val.fileName,
				words:    wordLen,
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	files := []string{
		"user_data_1.txt",
		"logs_2024_01.txt",
		"report_finance_q1.txt",
		"image_metadata_file0.txt",
		"image_metadata_file1.txt",
		"image_metadata_file.2txt",
		"image_metadata_file.t3xt",
		"image_metadata_file.tx4t",
		"image_metadata_file.txt5",
		"image_metadata_file.txt6",
		"image_metadata_file.txt7",
		"image_metadata_file.txt8",
		"image_metadata_file.txt9",
		"image_metadata_file.txt10",
		"image_metadata_file.txt11",
		"image_metadata_file.txt12",
		"image_metadata_file.txt13",
		"image_metadata_file.txt14",
		"image_metadata_file.txt15",
		"image_metadata_file.txt16",
		"image_metadata_file.txt17",
		"image_metadata_file.txt18",
		"image_metadata_file.txt20",
		"image_metadata_file.txt32",
		"backup_system_config.txt33",
	}
	ch := producer(files)
	results := make(chan Result, 5)
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go consumer(i, ctx, &wg, ch, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for val := range results {
		fmt.Println("My results", val)
	}
}
