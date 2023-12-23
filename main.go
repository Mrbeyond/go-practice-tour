package main

import (
	"fmt"
	"pract/files"
	"sync"
	"time"
)

// import "pract/advent"

// import "pract/stdin"

// import "pract/scrap"

func main() {

	// compr.CompressText()
	// compr.DecompressText("assets/test.txt")
	// compr.DoUndoCOmpress(compr.Base64Str)
	// apicall.TestMedilleryProjectsCount()
	// pdfgclearen.GeneratePdf()
	// xflag.Xflag()
	// xflag.RunFlags()
	// stdin.Simulator()
	// remotebase.Run()
	// scrap.Run()

	// advent.RunOne()
	// advent.RunTwo()
	// advent.RunFour()

	start := time.Now()
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go files.RunCurrent("/home/beyond/Desktop", ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for {
		log, valid := <-ch
		if valid {
			fmt.Println("\n\n", log, "\n\n ")
		} else {
			break
		}
	}

	fmt.Println(time.Since(start))
}
