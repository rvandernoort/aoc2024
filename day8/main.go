package main

import (
	"fmt"
	"time"

	"rovervandernoort.nl/framework"
)

func Part1(antennas map[string][]framework.Tuple, xMax int, yMax int) int {
	list := []framework.Tuple{} //todo make set
	for _, strAntennas := range antennas {
		for i := 0; i < len(strAntennas); i++ {
			for j := i + 1; j < len(strAntennas); j++ {
				xDiff := framework.AbsDiffInt(strAntennas[i].X, strAntennas[j].X)
				yDiff := framework.AbsDiffInt(strAntennas[i].Y, strAntennas[j].Y)
				if strAntennas[i].X < strAntennas[j].X && strAntennas[i].Y < strAntennas[j].Y {
					if strAntennas[i].X-xDiff >= 0 && strAntennas[i].Y-yDiff >= 0 {
						list = append(list, framework.Tuple{X: strAntennas[i].X - xDiff, Y: strAntennas[i].Y - yDiff})
					}
					if strAntennas[j].X+xDiff < xMax && strAntennas[j].Y+yDiff < yMax {
						list = append(list, framework.Tuple{X: strAntennas[j].X + xDiff, Y: strAntennas[j].Y + yDiff})
					}
				}
				if strAntennas[i].X > strAntennas[j].X && strAntennas[i].Y > strAntennas[j].Y {
					if strAntennas[i].X+xDiff < xMax && strAntennas[i].Y+yDiff < yMax {
						list = append(list, framework.Tuple{X: strAntennas[i].X + xDiff, Y: strAntennas[i].Y + yDiff})
					}
					if strAntennas[j].X-xDiff >= 0 && strAntennas[j].Y-yDiff >= 0 {
						list = append(list, framework.Tuple{X: strAntennas[j].X - xDiff, Y: strAntennas[j].Y - yDiff})
					}
				}
				if strAntennas[i].X < strAntennas[j].X && strAntennas[i].Y > strAntennas[j].Y {
					if strAntennas[i].X-xDiff >= 0 && strAntennas[i].Y+yDiff < yMax {
						list = append(list, framework.Tuple{X: strAntennas[i].X - xDiff, Y: strAntennas[i].Y + yDiff})
					}
					if strAntennas[j].X+xDiff < xMax && strAntennas[j].Y-yDiff >= 0 {
						list = append(list, framework.Tuple{X: strAntennas[j].X + xDiff, Y: strAntennas[j].Y - yDiff})
					}
				}
				if strAntennas[i].X > strAntennas[j].X && strAntennas[i].Y < strAntennas[j].Y {
					if strAntennas[i].X+xDiff < xMax && strAntennas[i].Y-yDiff >= 0 {
						list = append(list, framework.Tuple{X: strAntennas[i].X + xDiff, Y: strAntennas[i].Y - yDiff})
					}
					if strAntennas[j].X-xDiff >= 0 && strAntennas[j].Y+yDiff < yMax {
						list = append(list, framework.Tuple{X: strAntennas[j].X - xDiff, Y: strAntennas[j].Y + yDiff})
					}
				}
			}
		}
	}
	listMap := make(map[framework.Tuple]struct{})
	for _, tuple := range list {
		listMap[tuple] = struct{}{}
	}
	list = make([]framework.Tuple, 0, len(listMap))
	for tuple := range listMap {
		list = append(list, tuple)
	}
	return len(list)
}

func Part2(antennas map[string][]framework.Tuple, xMax int, yMax int) int {
	list := []framework.Tuple{}
	for _, strAntennas := range antennas {
		list = append(list, strAntennas...)
		for i := 0; i < len(strAntennas); i++ {
			for j := i + 1; j < len(strAntennas); j++ {
				xDiff := framework.AbsDiffInt(strAntennas[i].X, strAntennas[j].X)
				yDiff := framework.AbsDiffInt(strAntennas[i].Y, strAntennas[j].Y)
				if strAntennas[i].X < strAntennas[j].X && strAntennas[i].Y < strAntennas[j].Y {
					for k := 1; k <= xMax; k++ {
						if strAntennas[i].X-xDiff*k >= 0 && strAntennas[i].Y-yDiff*k >= 0 {
							list = append(list, framework.Tuple{X: strAntennas[i].X - xDiff*k, Y: strAntennas[i].Y - yDiff*k})
						}
						if strAntennas[j].X+xDiff*k < xMax && strAntennas[j].Y+yDiff*k < yMax {
							list = append(list, framework.Tuple{X: strAntennas[j].X + xDiff*k, Y: strAntennas[j].Y + yDiff*k})
						}
					}
				}
				if strAntennas[i].X > strAntennas[j].X && strAntennas[i].Y > strAntennas[j].Y {
					for k := 1; k <= xMax; k++ {
						if strAntennas[i].X+xDiff*k < xMax && strAntennas[i].Y+yDiff*k < yMax {
							list = append(list, framework.Tuple{X: strAntennas[i].X + xDiff*k, Y: strAntennas[i].Y + yDiff*k})
						}
						if strAntennas[j].X-xDiff*k >= 0 && strAntennas[j].Y-yDiff*k >= 0 {
							list = append(list, framework.Tuple{X: strAntennas[j].X - xDiff*k, Y: strAntennas[j].Y - yDiff*k})
						}
					}
				}
				if strAntennas[i].X < strAntennas[j].X && strAntennas[i].Y > strAntennas[j].Y {
					for k := 1; k <= xMax; k++ {
						if strAntennas[i].X-xDiff*k >= 0 && strAntennas[i].Y+yDiff*k < yMax {
							list = append(list, framework.Tuple{X: strAntennas[i].X - xDiff*k, Y: strAntennas[i].Y + yDiff*k})
						}
						if strAntennas[j].X+xDiff*k < xMax && strAntennas[j].Y-yDiff*k >= 0 {
							list = append(list, framework.Tuple{X: strAntennas[j].X + xDiff*k, Y: strAntennas[j].Y - yDiff*k})
						}
					}
				}
				if strAntennas[i].X > strAntennas[j].X && strAntennas[i].Y < strAntennas[j].Y {
					for k := 1; k <= xMax; k++ {
						if strAntennas[i].X+xDiff*k < xMax && strAntennas[i].Y-yDiff*k >= 0 {
							list = append(list, framework.Tuple{X: strAntennas[i].X + xDiff*k, Y: strAntennas[i].Y - yDiff*k})
						}
						if strAntennas[j].X-xDiff*k >= 0 && strAntennas[j].Y+yDiff*k < yMax {
							list = append(list, framework.Tuple{X: strAntennas[j].X - xDiff*k, Y: strAntennas[j].Y + yDiff*k})
						}
					}
				}
			}
		}
	}
	listMap := make(map[framework.Tuple]struct{})
	for _, tuple := range list {
		listMap[tuple] = struct{}{}
	}
	list = make([]framework.Tuple, 0, len(listMap))
	for tuple := range listMap {
		list = append(list, tuple)
	}
	return len(list)

}
func main() {
	lines := framework.ReadData("input.txt")
	antennas, xMax, yMax := framework.ParseAntennas(lines)

	fmt.Printf("-----\nPart 1\n")
	start := time.Now()
	product := Part1(antennas, xMax, yMax)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

	fmt.Printf("-----\nPart 2\n")
	start = time.Now()
	product = Part2(antennas, xMax, yMax)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

}
