package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
	"flag"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func processModule(i int) int{
	// divide i by 3
	m := i / 3.0

	//round down
	f := math.Floor(float64(m))
	m = int(f)

	if m > 0 {
		//subtract 2
		m = m - 2 
	}

	if m > 0 {
		// m is now module mass, but need to calc its own mass and add
		fmt.Printf("Module fuel: %d \n", m)
		m += processModule(m)
	}

	fmt.Printf("Total fuel: %d \n", m)
	if m < 0 {
		m = 0
	}
	return m
}

func main(){
	input := flag.Int("i", 0, "input")
	flag.Parse()

	fuel := 0
	if *input > 0 {
		fuel += processModule(*input)
	}else {
			file, err := os.Open("./input.txt")
			check(err)
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				// fmt.Println(scanner.Text())
				val, err := strconv.Atoi(scanner.Text())
				check(err)
				fuel += processModule(val)
			}
	}	

	fmt.Printf("Total fuel: %d", fuel)
}

