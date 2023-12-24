package test

import (
	"pract/files"
	"sync"
	"testing"
)

func BenchmarkFilepathWalDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		read := 0
		ch := make(chan int)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go files.CountFiles(".", ch, &wg)
		// go files.CountFiles("/home/beyond/Desktop", ch, &wg)
		go func() {
			wg.Wait()
			close(ch)
		}()

		for count := range ch {
			read += count
		}
	}
}

func BenchmarkManualRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		read := 0
		ch := make(chan int)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go files.RunCurrent(".", ch, &wg)
		// go files.RunCurrent("/home/beyond/Desktop", ch, &wg)
		go func() {
			wg.Wait()
			close(ch)
		}()

		for count := range ch {
			read += count
		}
	}
}
