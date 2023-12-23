// package advent

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// const red, green, blue = 12, 13, 14

// func computeLine(line string, ch chan<- int) {
// 	gameside := strings.Split(line, ":")[0]
// 	cubeside := strings.Split(line, ":")[1]

// 	bags := strings.Split(cubeside, ";")
// 	bagsLen := len(bags)
// 	invalid := false

// loopStart:
// 	for i := 0; i < bagsLen; i++ {
// 		cubes := strings.Split(strings.Trim(bags[i], " "), ",")
// 		cubesLen := len(cubes)
// 		for i := 0; i < cubesLen; i++ {
// 			valueAndColor := strings.Split(strings.Trim(cubes[i], " "), " ")
// 			value, _ := strconv.Atoi(valueAndColor[0])
// 			switch valueAndColor[1] {
// 			case "red":
// 				if value > red {
// 					invalid = true
// 					break loopStart
// 				}
// 			case "green":
// 				if value > red {
// 					invalid = true
// 					break loopStart
// 				}
// 			case "blue":
// 				if value > blue {
// 					invalid = true
// 					break loopStart
// 				}
// 			}
// 		}
// 	}

// 	if invalid {
// 		ch <- 0
// 		return
// 	}

// 	game, _ := strconv.Atoi(strings.Split(gameside, " ")[1])
// 	ch <- game
// }

// func posibleGames() (sum int) {
// 	ch := make(chan int)
// 	fmt.Println("Enter games \n ")
// 	scanner := bufio.NewScanner(os.Stdin)

// 	start := time.Now()
// 	go func() {
// 		for {
// 			scanner.Scan()
// 			game := scanner.Text()
// 			if len(game) == 0 {
// 				close(ch)
// 				break
// 			}
// 			go computeLine(game, ch)
// 		}
// 	}()

// 	for {
// 		res, valid := <-ch
// 		if !valid {
// 			break
// 		}
// 		sum += res
// 	}

// 	fmt.Println(sum, " Is Sum of games \n ")

// 	fmt.Println(time.Since(start))
// 	return sum
// }

// func RunThrree() {
// 	posibleGames()
// }

package advent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const red, green, blue = 12, 13, 14

var col []int

func computeLine(line string) (game int) {
	gameside := strings.Split(line, ":")[0]
	cubeside := strings.Trim((strings.Split(line, ":")[1]), " ")

	bags := strings.Split(cubeside, ";")
	bagsLen := len(bags)
	invalid := false

loopStart:
	for i := 0; i < bagsLen; i++ {
		cubes := strings.Split(strings.Trim(bags[i], " "), ",")
		cubesLen := len(cubes)
		for i := 0; i < cubesLen; i++ {
			valueAndColor := strings.Split(strings.Trim(cubes[i], " "), " ")
			value, _ := strconv.Atoi(valueAndColor[0])
			switch valueAndColor[1] {
			case "red":
				if value > red {
					invalid = true
					break loopStart
				}
			case "green":
				if value > green {
					invalid = true
					break loopStart
				}
			case "blue":
				if value > blue {
					invalid = true
					break loopStart
				}
			}
		}
	}

	if invalid {
		return
	}

	game, _ = strconv.Atoi(strings.Split(gameside, " ")[1])
	return game
}

func posibleGames() (sum int) {
	fmt.Println("Enter games \n ")
	scanner := bufio.NewScanner(os.Stdin)

	start := time.Now()
	for {
		scanner.Scan()
		game := scanner.Text()
		if len(game) == 0 {
			break
		}
		res := computeLine(game)
		col = append(col, res)
		sum += res
	}

	fmt.Println(sum, " Is Sum of games\n\n ")
	fmt.Println(col, "\n\n\n ")

	fmt.Println(time.Since(start))
	return sum
}

func RunThrree() {
	posibleGames()
}
