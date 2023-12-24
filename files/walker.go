package files

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

func print(a ...any) {
	fmt.Println("\n\nn", a, "\n\n\n ")
}

func CountFiles(path string, ch chan<- int, wg *sync.WaitGroup) {
	defer func() {
		// fmt.Println("Path: ", path, " \t ends here \n\n ")
		wg.Done()

	}()
	err := filepath.WalkDir(path, func(inpath string, d fs.DirEntry, err error) error {

		if err != nil && !errors.Is(err, os.ErrPermission) {
			fmt.Println("\n\n\n\n ERROR", err, "\n\n\n ")
			return err
		}

		if d.IsDir() {

		} else {
			ch <- 1
		}
		return nil
	})

	if err != nil {
		fmt.Println("\n\n\n\n ERROR", err, "\n\n\n ")
		return
	}
}
