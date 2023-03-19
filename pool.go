package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Jobz struct {
	ID int
}

type Resultz struct {
	Id string
}

var (
	jobCh = make(chan Jobz, 20)
	resCh = make(chan Resultz)
	rate  = make(map[int]int)
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func workerz(wg *sync.WaitGroup, workerIndex int) {
	for job := range jobCh {
		resCh <- Resultz{Id: fmt.Sprintf("job number => %d by worker %d", job, workerIndex)}
		mutex.Lock()
		_, ok := rate[workerIndex]
		if ok {
			rate[workerIndex]++
		} else {
			rate[workerIndex] = 1
		}
		mutex.Unlock()
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}

func allocateWorker(numOfWorker int) {
	for i := 0; i < numOfWorker; i++ {
		wg.Add(1)
		go workerz(&wg, i+1)
	}
	wg.Wait()
	close(resCh)
}

func allocateJob(numberOfJobs int) {
	for i := 0; i < numberOfJobs; i++ {
		jobCh <- Jobz{ID: i + 1}
	}
	close(jobCh)
}

func examiner(done chan bool) {
	for res := range resCh {
		println(fmt.Sprintf("Resultz for %s \n", res))
	}
	done <- true
}

func smain() {
	startTime := time.Now()
	done := make(chan bool)

	simulateBuffer()
	numOfWorker := 50
	numberOfJobs := 2000

	go allocateJob(numberOfJobs)
	go examiner(done)
	allocateWorker(numOfWorker)

	<-done
	// CreateCSV()

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Milliseconds(), "Milliseconds")
	fmt.Println(rate)
}

func simulateBuffer() {
	mainChunk := make([][]int, 3010)
	tempChunk := make([]int, 3)
	for i := 0; i < 9030; i++ {
		index := i % 3
		tempChunk[index] = i + 1
		if index == 2 {
			mainChunk[i/3] = tempChunk
			tempChunk = make([]int, 3)
		}
	}

	var buf bytes.Buffer
	var err error
	writer := csv.NewWriter(&buf)

	for _, record := range mainChunk {
		intToStringRow := make([]string, len(record))
		for r, value := range record {
			intToStringRow[r] = fmt.Sprintf("%d", value)
		}
		if err = writer.Write(intToStringRow); err != nil {
			panic(err)
		}
	}
	writer.Flush()

	reader := csv.NewReader(&buf)
	// Set the expected number of fields per record
	reader.FieldsPerRecord = 3
	buffer := make([][]string, 1000) // buffer for 1000 records
	rev := 1

	for {
		fmt.Print("\n\n Revolution ", rev, " \n\n")

		// Read the next chunk of records
		for i := 0; i < 1000; i++ {
			record, err := reader.Read()
			// Check for end of file
			if err == io.EOF {
				break
			}
			// Check for other errors
			if err != nil {
				panic(err)
			}
			// Add the record to the buffer
			buffer[i] = record
		}

		println("Statge two")
		// Validate the records in the buffer
		for i := 0; i < len(buffer); i++ {
			record := buffer[i]
			// Check for end of buffer
			if record == nil {
				break
			}
			// Validate the record
			if len(record) != reader.FieldsPerRecord {
				fmt.Println("Record has incorrect number of fields")
				return
			}
		}

		fmt.Println("\n\n Burffer for rev ", rev, "\n\n", buffer, "\n\n ")

		// Check for end of file
		if len(buffer[0]) == 0 {
			println("\n Break block \n ")
			fmt.Println(buffer)
			break
		}
		// Clear the buffer
		buffer = make([][]string, 1000)
	}
}

// Create  a csv file
func CreateCSV() {
	type OHLC struct {
		UNIX   uint64
		SYMBOL string
		OPEN   float64
		HIGH   float64
		LOW    float64
		CLOSE  float64
	}

	fields := []OHLC{
		{UNIX: 1644719700000, SYMBOL: "BTCUSDT", OPEN: 42123.29000000, HIGH: 42148.32000000, LOW: 42120.82000000, CLOSE: 42146.06000000},
		{UNIX: 1644719640000, SYMBOL: "BTCUSDT", OPEN: 42113.08000000, HIGH: 42126.32000000, LOW: 42113.07000000, CLOSE: 42123.30000000},
		{UNIX: 1644719580000, SYMBOL: "BTCUSDT", OPEN: 42120.80000000, HIGH: 42130.23000000, LOW: 42111.01000000, CLOSE: 42113.07000000},
		{UNIX: 1644719520000, SYMBOL: "BTCUSDT", OPEN: 42114.47000000, HIGH: 42123.31000000, LOW: 42102.22000000, CLOSE: 42120.80000000},
		{UNIX: 1644719460000, SYMBOL: "BTCUSDT", OPEN: 42148.23000000, HIGH: 42148.24000000, LOW: 42114.04000000, CLOSE: 42114.48000000},
	}

	max := 50000
	// data := make([][]string, max)

	// var buf bytes.Buffer
	var err error
	file, err := os.OpenFile("test.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// var header = []string{"UNIX", "SYMBOL", "OPEN", "HIGH", "LOW", "CLOSE"}
	// if err = writer.Write(header); err != nil {
	// 	panic(err)
	// }

	for i := 0; i < max/5; i++ {
		// domain := i * len(fields)
		for _, field := range fields {
			record := []string{
				fmt.Sprintf("%d", field.UNIX),
				field.SYMBOL,
				fmt.Sprintf("%f", field.OPEN),
				fmt.Sprintf("%f", field.HIGH),
				fmt.Sprintf("%f", field.LOW),
				fmt.Sprintf("%f", field.CLOSE),
			}

			if err = writer.Write(record); err != nil {
				panic(err)
			}
		}
	}

	// fmt.Println(data[0:5], "\n\n" /* data[4989:5000],*/, "\n\n len=> ", len(data), "\n\n ")

}
