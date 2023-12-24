package files

import (
	"sync"
	"testing"
)

func BenchmarkFilepathWalDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		read := 0
		ch := make(chan int)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go CountFiles(".", ch, &wg)
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
		go RunCurrent(".", ch, &wg)
		go func() {
			wg.Wait()
			close(ch)
		}()

		for count := range ch {
			read += count
		}
	}
}
