package scanner

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	slCh = make(chan []string)
	rsCh = make(chan string, 20)
	// done        = make(chan bool)
	numOfWorker = 80
	wx          sync.WaitGroup
)

func Slice2h(count int) (record [][]string) {
	record = make([][]string, 0)
	for i := 1; i <= count; i++ {
		record = append(record, []string{fmt.Sprintf("%d", i)})
	}
	return
}

func FillSlCh(slice [][]string) {
	file, _ := os.Open("test2.csv")
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		slCh <- row
	}
	// for _, row := range slice {
	// 	slCh <- row
	// }
	println("One")
	close(slCh)
}

func procRes(done chan bool) {
	resValue := ""
	for value := range rsCh {
		resValue += value + "\n"
	}
	fmt.Println("\n val =>", resValue)
	fmt.Println("\n", len(strings.Split(strings.Trim(resValue, "\n"), "\n")), "\n ")
}

func slChAction(wg *sync.WaitGroup) {
	defer wg.Done()
	for record := range slCh {
		rsCh <- record[0]
		println(record[0])
		time.Sleep(50 * time.Millisecond)
	}
}
func slChPool(numOfWorker int) {
	for i := 0; i < numOfWorker; i++ {
		// println(i)
		wx.Add(1)
		go slChAction(&wx)
	}
}

func scan() {
	startTime := time.Now()
	slice := Slice2h(200)
	// done := make(chan bool)
	//Job
	// go procRes(done)
	go FillSlCh(slice)
	slChPool(numOfWorker)
	resValue := ""
	go func() {
		for value := range rsCh {
			resValue += value + "\n"
		}
	}()
	// <-done
	wx.Wait()

	fmt.Println("\n", resValue, "\n ")
	fmt.Println("\n", len(strings.Split(strings.Trim(resValue, "\n"), "\n")), "\n ")
	// <-done
	// close(done)

	numCores := runtime.NumCPU()
	numWorkers := numCores
	fmt.Printf("Number of worker pool threads: %d\n", numWorkers)
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Milliseconds(), "Milliseconds")
}
