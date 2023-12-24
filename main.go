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
	read := 0
	ch := make(chan int)
	wg := sync.WaitGroup{}
	// chw := make(chan int)
	// wgw := sync.WaitGroup{}
	wg.Add(1)
	// wgw.Add(1)

	// go files.RunCurrent("/home/beyond/Desktop", ch, &wg)
	// go files.RunCurrent(".", ch, &wg)
	go files.CountFiles("/home/beyond/Desktop", ch, &wg)
	// go files.CountFiles(".", ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	// go func() {
	// 	wgw.Wait()
	// 	close(chw)
	// }()

	for {
		// select {
		// case
		count, valid := <-ch
		if valid {

			read += count
		} else {
			fmt.Println("Invalid from logger")
			break
		}
		// case
		// counter, valid := <-chw
		// if valid {

		// 	read += counter
		// 	fmt.Println("\n\n from counter", read)
		// } else {
		// 	fmt.Println("Invalid from counter")
		// 	break
		// }
		// default:
		// 	fmt.Println("Invalid from default")
		// 	return

	}

	// }

	// for count := range chw {
	// 	read += count
	// 	fmt.Println("\n\n from counter", read)

	// }

	fmt.Println("\n\n Files read : ", read)
	fmt.Println(time.Since(start))
}
