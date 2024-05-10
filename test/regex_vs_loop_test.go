package test

import (
	"bufio"
	"flag"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

var inputFlag = flag.String("input", "", "Enter value")
var count = 0

func useRegex(str string) (sum int, err error) {
	flag.Parse()
	count = 0
	var rr *strings.Reader
	if *inputFlag != "" {
		rr = strings.NewReader(*inputFlag)
	} else {
		rr = strings.NewReader(str)
	}
	rg := regexp.MustCompile(`\d`)
	scanner := bufio.NewScanner(rr)

	for {
		scanner.Scan()
		str := scanner.Text()
		if len(str) == 0 {
			break
		}
		matches := rg.FindAllString(str, -1)
		firstDigit, lastDigit := "", ""
		switch len(matches) {
		case 0:
		case 1:
			firstDigit = matches[0]
			lastDigit = firstDigit
		default:
			firstDigit = matches[0]
			lastDigit = matches[len(matches)-1]
		}

		res, _ := strconv.Atoi(firstDigit + lastDigit)
		sum += res
		count++
		// fmt.Println("Text is ", text)
	}
	// fmt.Println("Cunt is ", count)
	// fmt.Println(sum, " Is sum Reg")

	return sum, err
}

func useLoop(str string) (sum int) {
	var rr *strings.Reader
	if *inputFlag != "" {
		rr = strings.NewReader(*inputFlag)
	} else {
		rr = strings.NewReader(str)
	}
	scanner := bufio.NewScanner(rr)

	for {
		scanner.Scan()
		str := scanner.Text()
		size := len(str)
		if size == 0 {
			break
		}
		start, end := "", ""
		for i := 0; i < size; i++ {
			if unicode.IsDigit(rune(str[i])) {
				start = string(str[i])
				break
			}
		}

		for i := size - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(str[i])) {
				end = string(str[i])
				break
			}
		}

		res, _ := strconv.Atoi(start + end)
		sum += res
	}

	// fmt.Println(sum, " Is sum loop")

	return sum
}

const str = "abweqjkjqw qwjejdic123def456xyz789kokok jkwekdoqwkdowkdoqwk"

func BenchmarkLoopForRegexVsLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useLoop(str)
	}
}

func BenchmarkRegexForLoopVsRegex(b *testing.B) {
	// fmt.Println(count)
	for i := 0; i < b.N; i++ {
		useRegex(str)
	}
}
