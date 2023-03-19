package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world")
}

func numbers() {
	for i := 0; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d => num \n", i)
	}
}
func alpha() {
	for i := 'a'; i < 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c => char \n", i)
	}
}

func mains() {
	// go numbers()
	// go alpha()
	// time.Sleep(3000 * time.Millisecond)
	// fmt.Println("This the main entry ends")

	lt := time.Now().Nanosecond()
	fmt.Println(lt)
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
	fmt.Println(time.Now().Nanosecond() - lt)

	lt = time.Now().Nanosecond()
	fmt.Println(lt, "sec")
	sqrch2 := make(chan int)
	cubech2 := make(chan int)
	go calcSquares2(number, sqrch2)
	go calcCubes2(number, cubech2)
	squares, cubes = <-sqrch2, <-cubech2
	fmt.Println("Final output", squares+cubes)
	fmt.Println(time.Now().Nanosecond() - lt)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}
func calcSquares2(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes2(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}
