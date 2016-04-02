package main

import "fmt"

// import "strconv"
import "strings"
import "regexp"

func main() {
	zeroMatrix := make([][]int, 17, 17)
	str := "Babaji@27AE60.com"
	normalizedEmail := strings.ToLower(regexp.MustCompile("[^a-zA-Z0-9]").ReplaceAllString(email, ""))

	// initialze
	for i := 0; i < 10; i++ {
		zeroMatrix[i] = make([]int, 17, 17)
	}

	// initial seed
	fmt.Println(hexCode, normalizedEmail)

	/*
		// print initial state
		for i := 0; i < 10; i++ {
			fmt.Println(zeroMatrix[i])
		}

		fmt.Println("-------------------------------------")

		for z := 0; z < 10; z++ {
			result := make([][]int, 10, 10)

			for i := 1; i < 9; i++ {
				result[i] = make([]int, 10, 10)

				for j := 1; j < 9; j++ {
					neighbourCount := 0

					if zeroMatrix[i-1][j-1] == 1 {
						neighbourCount++
					}
					if zeroMatrix[i-1][j] == 1 {
						neighbourCount++
					}
					if zeroMatrix[i-1][j+1] == 1 {
						neighbourCount++
					}
					if zeroMatrix[i][j-1] == 1 {
						neighbourCount++
					}
					if zeroMatrix[i][j+1] == 1 {
						neighbourCount++
					}
					if zeroMatrix[i+1][j-1] == 1 {
						neighbourCount++
					}
					if zeroMatrix[i+1][j] == 1 {
						neighbourCount++
					}
					if zeroMatrix[i+1][j+1] == 1 {
						neighbourCount++
					}

					if zeroMatrix[i][j] == 1 {
						if neighbourCount == 2 || neighbourCount == 3 {
							result[i][j] = 1
						} else {
							result[i][j] = 0
						}
					} else {
						if neighbourCount == 3 {
							result[i][j] = 1
						}
					}

					for x := 0; x < 10; x++ {
						fmt.Println(result[x])
					}
					fmt.Println("-------------------------------------")
				}
			}

			zeroMatrix = result
		}
	*/
}
