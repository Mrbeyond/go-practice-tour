package xflag

import (
	"flag"
	"fmt"
	"os"
)

func Xflag() {
	fmt.Println(os.Getwd())
	fmt.Println(os.Args)
}

func RunFlags() {
	// Defining a string flag
	strFlag := flag.String("language", "Golang", "Golang is the awesome google language")
	// Defining an integer flag
	intFlag := flag.Int("downloads", 1000000, "Number of times Go has been downloaded")
	// Defining a boolean flag
	boolFlag := flag.Bool("isAwesome", true, "Yes! Go is awesome")

	// Call flag.Parse() to parse the command-line flags
	flag.Parse()
	// Log the flags to the terminal
	fmt.Printf("String flag %v \n", *strFlag)
	fmt.Println("Integer flag ", *intFlag)
	fmt.Println("Boolean flag ", *boolFlag)
}
