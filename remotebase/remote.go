package remotebase

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
)

/*
 * Complete the 'writeToFile' function below.
 *
 * The function accepts following parameters:
 *  1. chan []byte bytesChannel
 *  2. chan bool doneChannel
 *  3. chan error errChannel
 */

func writeToFile(bytesChannel chan []byte, doneChannel chan bool, errChannel chan error) {

	file, err := os.Create(filename)
	errChannel <- err
	if err != nil {
		return
	}
	defer file.Close()

	go func() {
		for {
			select {
			case <-doneChannel:
				close(bytesChannel)
				close(errChannel)
				close(doneChannel)
				return
			default:
			}
		}
	}()

	for fileByte := range bytesChannel {
		_, err := file.Write(fileByte)
		errChannel <- err
		if err != nil {
			return
		}
	}

}

func Run() {
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	// inputArrayCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// checkError(err)

	inputArray := []string{"append", "inputArray", "inputArrayCount"}

	// for i := 0; i < int(inputArrayCount); i++ {
	// 	inputArrayItem := readLine(reader)
	// 	inputArray = append(inputArray, inputArrayItem)
	// }

	bytesChannel, doneChannel, errChannel := make(chan []byte), make(chan bool), make(chan error)
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	allocBefore := ms.Alloc
	fmt.Println(allocBefore, " >> allocBefore \n ")
	go writeToFile(bytesChannel, doneChannel, errChannel)
	err := <-errChannel
	if err != nil {
		panic(err)
	}
	for _, b := range inputArray {
		bytesChannel <- []byte(b)
		err := <-errChannel
		if err != nil {
			// fmt.Fprintf(writer, "Critical error: %s", err.Error())
			panic(err)
		}
	}
	doneChannel <- true
	runtime.ReadMemStats(&ms)
	allocAfter := ms.Alloc
	log.Println(allocAfter, " >> allocAfter \n ")
	fmt.Printf("Total memory allocated: %d bytes\n", allocAfter-allocBefore)
	if allocAfter-allocBefore > 10000 {
		log.Fatal("Too much memory allocated, maximum 10000 bytes needed")
	} else {
		b, err := ioutil.ReadFile(filename)
		if err == nil {
			log.Fatal(string(b))
		} else {
			log.Fatalf("Critical error: %s", err.Error())
		}
	}
}

const filename = "output"

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
