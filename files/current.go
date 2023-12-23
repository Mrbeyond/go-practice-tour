package files

import (
	"fmt"
	"os"
)

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

var read = 0

func RunCurrent(path string, ch chan<- string) {
	defer close(ch)
	entries, err := os.ReadDir(path)
	fmt.Println(len(entries), "\t files", "\t\t path -> \t", path, "\n\n ")
	if os.IsPermission(err) {
		fmt.Sprintln(err)
	} else {
		panicError(err)
	}

	for _, file := range entries {
		// fmt.Println("", file.Name(), file.IsDir(), "\n\n ")

		if file.IsDir() {
			go RunCurrent(path+"/"+file.Name(), ch)
		} else {
			read++
			stat, err := os.Stat(path + "/" + file.Name())
			// panicError(err)
			if err != nil {
				fmt.Println("\n\n\n\n", err, "\n\n\n\n\n\n ")
			} else {

				ch <- fmt.Sprintf(" %s \t\t\t %v \t\t\t %s", file.Name(), stat.Size(), stat.Mode())
			}
		}
	}
	println("\n\n read ", read, "files\n ")
}
