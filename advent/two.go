package advent

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var words = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func convertWordsTofigure(str string) string {
	position := 0

StartSide:
	for position < len(str) {
		for key, value := range words {
			if strings.HasPrefix(str[position:], key) {
				str = strings.Replace(str, key, value, 1)
				break StartSide
			} else if unicode.IsDigit(rune(str[position])) {
				break StartSide
			}
		}
		position++
	}

	position = len(str) - 1

OuterLoop:
	for position >= 0 {
		for key, value := range words {
			if strings.HasSuffix(str[:position+1], key) {
				str = str[:position-len(key)] + strings.Replace(str[position-len(key):], key, value, 1)
				break OuterLoop
			} else if unicode.IsDigit(rune(str[position])) {
				break OuterLoop
			}
		}
		position--
	}
	return str
}

func caliberateWords() (sum int, err error) {
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
		str = convertWordsTofigure(str)
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

func RunTwo() {
	caliberateWords()
}
