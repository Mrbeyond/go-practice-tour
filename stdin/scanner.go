package stdin

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func panicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
func UseFmtScan() {
	fmt.Println("FMT SCAN: Enter your first, second and  last name")
	var first, second, last string
	n, err := fmt.Scan(&first, &second, &last)
	panicError(err)
	fmt.Printf(" Number of arguements %d n", n)
	fmt.Printf("Your first name is %s \n Your second name is %s \n Your last name is %s \n",
		first, second, last,
	)
}

func UseFmtScanLine() {
	fmt.Println("FMT SCAN SINGLE LINE: Enter your first, second and  last name")
	var first, second, last string
	n, err := fmt.Scanln(&first, &second, &last)
	panicError(err)
	fmt.Printf(" Number of arguements %d n", n)
	fmt.Printf("Your first name is %s \n Your second name is %s \n Your last name is %s \n",
		first, second, last,
	)
}

func UseBufioReadString() {
	fmt.Println("Bufio Readstring Single: Enter your name")
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	panicError(err)
	fmt.Println("Your name is ", value)
}

func UseBufioScan() {
	fmt.Println("Bufio Scan Single: Enter your name")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	panicError(scanner.Err())

	fmt.Println("Your name is ", scanner.Text())
}

func UseBufioReadStringMultiple() {
	fmt.Println("Bufio Readstring Multiple: Enter multiple values ")
	reader := bufio.NewReader(os.Stdin)
	for {
		value, err := reader.ReadString('\n')
		panicError(err)
		fmt.Println("You entered ", value)
		if len(strings.TrimSpace(value)) == 0 {
			break
		}
	}
}

func UseBufioScanMultiple() {
	fmt.Println("Bufio Scan Multiple: Enter multiple values ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
		fmt.Println("Your name is ", text)
	}
	panicError(scanner.Err())
}

func Simulator() {
	UseFmtScan()
	UseFmtScanLine()
	UseBufioReadString()
	UseBufioScan()
	UseBufioReadStringMultiple()
	UseBufioScanMultiple()
}
