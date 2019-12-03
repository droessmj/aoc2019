package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"strconv"
)

const mark = "x"
const intersection = "i"
const start = "s"

type route struct {
	moves	[]move
}
type move struct {
	direction string
	distance  int
}
type junction struct {
	x 	int
	y  	int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseMovementsFromString(s string) route {
	moves := make([]move, 0)
	movements := strings.Split(s, ",")

	//parse the direction and the distance from each input
	for _, m  := range movements{
		distance, e := strconv.Atoi(string(m[1:]))
		check(e)
		move := move {
			direction: string(m[0]),
			distance: distance,
		}
		moves = append(moves, move)
	}

	fmt.Println(moves)
	return route { moves: moves }
}

func addRoute(r route, t [][]string, center junction) ([][]string, []junction) {
	curX := center.x
	curY := center.y	

	junctions := make([]junction,0)
	for _, m := range r.moves {
		switch m.direction {
			case "U":
				for i := 0; i < m.distance; i++ {
					curY+=1
					//mark the moves if not marked. If marked, mark as an intersection
					if t[curX][curY] == mark {
						junctions = append(junctions, junction{ x: curX, y: curY})
						t[curX][curY] = intersection
					} else {
						t[curX][curY] = mark
					}
				}
			case "D":
					curY-=1
					//mark the moves if not marked. If marked, mark as an intersection
					if t[curX][curY] == mark {
						junctions = append(junctions, junction{ x: curX, y: curY})
						t[curX][curY] = intersection
					} else {
						t[curX][curY] = mark
					}
			case "L":
					curX-=1
					//mark the moves if not marked. If marked, mark as an intersection
					if t[curX][curY] == mark {
						junctions = append(junctions, junction{ x: curX, y: curY})
						t[curX][curY] = intersection
					} else {
						t[curX][curY] = mark
					}
			case "R":
					curX+=1
					//mark the moves if not marked. If marked, mark as an intersection
					if t[curX][curY] == mark {
						junctions = append(junctions, junction{ x: curX, y: curY})
						t[curX][curY] = intersection
					} else {
						t[curX][curY] = mark
					}
			default:
				panic("Whyyyy222")
		}
	}

	return t, junctions
}

func calculateIntersection(j []junction, center junction) int{
	result := math.MaxInt32

	return result
}

func main() {

	file, err := os.Open("./input4.txt")
	check(err)
	defer file.Close()

	routes := make([]route,2)
	scanner := bufio.NewScanner(file)

	// build board and feed it into the moves builder
	maxR, maxL, maxU, maxD := 0, 0, 0, 0

	i := 0
	for scanner.Scan() {
		routes[i] = parseMovementsFromString(scanner.Text())
		for _, m := range routes[i].moves {
			switch m.direction {
				case "U":
					maxU += m.distance
				case "D":
					maxD += m.distance
				case "L":
					maxL += m.distance
				case "R":
					maxR += m.distance
				default:
					panic("Whyyyy")
			}
		}
		i++
	}

	// +1 because these are all offsets from center
	width := maxR + maxL + 1
	height := maxU + maxD + 1
	fmt.Printf("\nmaxU: %d, maxD: %d, maxL: %d, maxR: %d, width: %d, height: %d\n", maxU, maxD, maxL, maxR, width, height)

	t := make([][]string, width)
	for i := 0; i < width; i++ {
		t[i] = make([]string, height)
	}

	//mark the "center" junction
	center := junction {
		x: width - maxR,
		y: height - maxU,
	}
	fmt.Println(t)
	fmt.Printf("\ncenterX: %d, centerY: %d\n", center.x, center.y)

	t[center.x][center.y] = start
	center = junction{ x: center.x, y: center.y}


	junctions := make([]junction,0)
	j := make([]junction,0)
	for _, route := range routes{
		t, j = addRoute(route, t, center)
		junctions = append(junctions, j...)
	}
	for i := 0; i < width; i++ {
		fmt.Println(t[i])
	}
	//need to find the junctions
	
	fmt.Printf("\njunctions tracked: %d\n", len(junctions))
	fmt.Println(junctions)

	result := calculateIntersection(junctions, center)

	fmt.Printf("Shortest Distance: %d", result)
}
