package files

import (
	"fmt"
	"os"
	"sync"
)

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

func RunCurrent(path string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	entries, err := os.ReadDir(path)
	if os.IsPermission(err) {
		fmt.Sprintln(err)
	} else {
		panicError(err)
	}

	for _, file := range entries {
		if file.IsDir() {
			wg.Add(1)
			go RunCurrent(path+"/"+file.Name(), ch, wg)
		} else {
			// _, err := os.Stat(path + "/" + file.Name())
			// panicError(err)
			// if err != nil {
			// fmt.Println("\n\n\n\n", err, "\n\n\n\n\n\n ")
			// } else {

			ch <- 1
			// }
		}
	}
}
