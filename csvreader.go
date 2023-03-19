package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	// _ "github.com/go-sql-driver/mysql"
)

type Row struct {
	ID   int
	Name string
	Age  int
}

const batchSize = 100000

func ss() {
	startTime := time.Now()

	// Create database connection
	db, err := sql.Open("mysql", "user:password@tcp(host:port)/database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Open CSV file
	file, err := os.Open("large_file.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create buffered reader
	reader := bufio.NewReader(file)

	// Create worker pool
	workerCount := 4 // adjust the number of workers to your machine's capabilities
	jobs := make(chan []string, workerCount)
	results := make(chan error, workerCount)
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go workers(db, jobs, results, &wg)
	}

	// Process CSV data
	var lineCount int
	var chunk []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		lineCount++
		chunk = append(chunk, line)

		// Check chunk size
		if len(chunk) == batchSize {
			// Add job to worker pool
			jobs <- chunk
			chunk = make([]string, 0, batchSize)
		}
	}

	// Add any remaining lines to worker pool
	if len(chunk) > 0 {
		jobs <- chunk
	}

	// Close jobs channel to signal workers to stop
	close(jobs)

	// Wait for workers to finish
	wg.Wait()

	// Check for any errors
	for i := 0; i < workerCount; i++ {
		if err := <-results; err != nil {
			panic(err)
		}
	}

	fmt.Println("Time taken:", time.Since(startTime))
	fmt.Println("Total lines processed:", lineCount)
}

func workers(db *sql.DB, jobs <-chan []string, results chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	for chunk := range jobs {
		// Create batch slice
		batch := make([]Row, 0, batchSize)

		// Parse CSV data
		reader := csv.NewReader(strings.NewReader(strings.Join(chunk, "")))
		for {
			fields, err := reader.Read()
			if err != nil {
				break
			}

			// Create row
			id, err := strconv.Atoi(fields[0])
			if err != nil {
				results <- err
				return
			}
			age, err := strconv.Atoi(fields[2])
			if err != nil {
				results <- err
				return
			}
			row := Row{ID: id, Name: fields[1], Age: age}

			// Add row to batch
			batch = append(batch, row)

			// Check batch size
			if len(batch) == batchSize {
				// Insert batch into database
				err := insertBatch(db, batch)
				if err != nil {
					results <- err
					return
				}
			}
		}
	}
}

func insertBatch(db *sql.DB, batch []Row) error {
	panic("unimplemented")
}
