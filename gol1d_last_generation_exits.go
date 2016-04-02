package main

import "fmt"

func main() {
	birth := make([]int, 20, 20)
	startPoints := []int{2, 9, 18}

	for i := 0; i < len(startPoints); i++ {
		birth[startPoints[i]] = 1
	}

	fmt.Println(birth)

	tick := true

	for z := 0; z < 20; z++ {
		var result []int

		if tick != true {
			result = make([]int, 20, 20)
		} else {
			result = birth
		}

		for i := 1; i < 19; i++ {
			amount := 0
			value := 0

			if birth[i+1] == 1 {
				amount++
			}
			if birth[i-1] == 1 {
				amount++
			}
			if amount == 1 {
				value = 1
			}

			result[i] = value
		}

		fmt.Println(result)

		birth = result
	}
}
