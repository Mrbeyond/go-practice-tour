package advent

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

var count = 0

func useRegex() (sum int, err error) {
	count = 0

	rg := regexp.MustCompile(`\d`)
	fmt.Println("Bufio Regex values ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		str := scanner.Text()
		if len(str) == 0 {
			break
		}
		matches := rg.FindAllString(str, -1)
		fmt.Println("\n", matches, " matches \n ")
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
		// fmt.Println("Text is ", text)
	}
	// fmt.Println("Cunt is ", count)
	fmt.Println(sum, " Is Sum Reg")
	fmt.Println(count, " Is Count Reg")

	return sum, err
}

func useLoop() (sum int) {
	count = 0

	fmt.Println("Bufio Loop values ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		str := scanner.Text()
		if len(str) == 0 {
			break
		}
		start, end := "", ""
		for i := 0; i < len(str); i++ {
			if unicode.IsDigit(rune(str[i])) {
				start = string(str[i])
				break
			}
		}

		for i := len(str) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(str[i])) {
				end = string(str[i])
				break
			}
		}

		fmt.Println("\n", str, " start => end :", start, " => ", end, "\n ")
		res, _ := strconv.Atoi(start + end)
		sum += res
		count++
	}

	fmt.Println(count, " Is Count Loop")
	fmt.Println(sum, " Is Sum Loop")

	return sum
}

func RunOne() {
	useLoop()
	useRegex()
}
