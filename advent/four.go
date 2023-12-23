package advent

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var greenReg = regexp.MustCompile(`\d+\sgreen`)
var redReg = regexp.MustCompile(`\d+\sred`)
var blueReg = regexp.MustCompile(`\d+\sblue`)
var intReg = regexp.MustCompile(`\d+`)

func findMax(values []string) (max int) {

	min := math.MinInt
	for _, value := range values {
		num, _ := strconv.Atoi(value)
		if num > min {
			min = num
			max = num
		}
	}
	return
}

func computeFewestCubes(line string) int {

	reds := strings.Join(redReg.FindAllString(line, -1), ",")
	greens := strings.Join(greenReg.FindAllString(line, -1), ",")
	blues := strings.Join(blueReg.FindAllString(line, -1), ",")

	redValues := intReg.FindAllString(reds, -1)
	greenValues := intReg.FindAllString(greens, -1)
	blueValues := intReg.FindAllString(blues, -1)

	return findMax(redValues) * findMax(greenValues) * findMax(blueValues)
}

func fewestCube() (sum int) {
	fmt.Println("Enter games \n ")
	scanner := bufio.NewScanner(os.Stdin)

	start := time.Now()
	for {
		scanner.Scan()
		game := scanner.Text()
		if len(game) == 0 {
			break
		}
		res := computeFewestCubes(game)
		col = append(col, res)
		sum += res
	}

	fmt.Println(sum, " Is Sum of games\n\n ")
	fmt.Println(col, "\n\n\n ")

	fmt.Println(time.Since(start))
	return sum
}

func RunFour() {
	fewestCube()
}
