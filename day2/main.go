package main

import (
	"fmt"
	"flag"
	"strings"
	"strconv"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func overflowIndex(length int, value int) int {
	if value > length{
		value = value - length
	}
	return value
}

func adder(t []int, i int, j int) int{
	return t[i] + t[j]
}

func multiplier(t []int, i int, j int) int{
	return t[i] * t[j]
}

func calculateOutput (t []int) int {

	fmt.Println("initial: ", t)
	original := make([]int, len(t))
	copy(original, t[:])

	// seeds -- all between 0 and 99
	notFound := true
	pos1 := 0

	for notFound == true && pos1 < 100{
			pos2 := 0
			for pos2 < 100 {
					copy(t, original[:])
					t[1] = pos1
					t[2] = pos2
					curIdx := 0
					opCode := t[curIdx]

					for opCode != 99 {
						idxToSet := t[overflowIndex(len(t), curIdx + 3)]
						valToSet := 0

						i := t[overflowIndex(len(t), curIdx+1)]
						j := t[overflowIndex(len(t), curIdx+2)]

						switch opCode {
							case 1:
								valToSet = adder(t, i, j)
							case 2: 
								valToSet = multiplier(t, i, j)	
							case 99:
								break
							default:
								break
						}
						
						if idxToSet > len(t){
							break
						}

						t[idxToSet] = valToSet
						curIdx = overflowIndex(len(t), curIdx+4)	
						opCode = t[curIdx]
					}
					// target == 19690720
					if t[0] == 19690720{
						fmt.Println(t)
						fmt.Printf("\n1: %d, 2: %d\n", pos1, pos2)
						fmt.Printf("\nanswer: %d", (100*pos1)+pos2)
						return t[0]
					}
					pos2++
			}
			pos1++
	}

	return t[0]
}

func processInput(s string) int{

	// split the input string into rows of 4 or input 99
	split := strings.Split(s, ",")
	input := make([]int, len(split))
	pos := 0
	for _, i := range split {
		j, err := strconv.Atoi(string(i))
		check(err)
		input[pos] = j
		pos++
	}

	return calculateOutput(input)
}

func main(){
	input := flag.String("i", "", "input")
	flag.Parse()

	val := 0
	if len(*input) > 0{
		val = processInput(*input)
	}else {
			s := "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,6,19,1,9,19,23,2,23,10,27,1,27,5,31,1,31,6,35,1,6,35,39,2,39,13,43,1,9,43,47,2,9,47,51,1,51,6,55,2,55,10,59,1,59,5,63,2,10,63,67,2,9,67,71,1,71,5,75,2,10,75,79,1,79,6,83,2,10,83,87,1,5,87,91,2,9,91,95,1,95,5,99,1,99,2,103,1,103,13,0,99,2,14,0,0"

			//file, err := os.Open("./input.txt")
			//check(err)
			//defer file.Close()

			//scanner := bufio.NewScanner(file)
			//for scanner.Scan() {
			//	val = processInput(scanner.Text())
			//}
			val = processInput(s)
	}	

	fmt.Printf("\nOutcome: %d", val)
}

