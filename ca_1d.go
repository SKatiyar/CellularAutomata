package main

import "sync"
import "fmt"

type cell struct {
	x int
	l int
}

func main() {
	previousState := make([]int, 10, 10)
	initialSeeds := []int{5}

	for i := 0; i < len(initialSeeds); i++ {
		previousState[initialSeeds[i]] = 1
	}

	stateChan := make(chan cell)

	for i := 0; i < 10; i++ {
		var waitGroup sync.WaitGroup

		fmt.Println(previousState, "======")

		for j := 0; j < len(previousState); j++ {
			neighbours := make([]int, 0)
			if j == 0 {
				neighbours = append(neighbours, previousState[j+1])
			} else if j == len(previousState)-1 {
				neighbours = append(neighbours, previousState[j-1])
			} else {
				neighbours = append(neighbours, previousState[j+1])
				neighbours = append(neighbours, previousState[j-1])
			}

			waitGroup.Add(1)
			go computeCell(stateChan, cell{x: j, l: previousState[j]}, neighbours...)
		}

		newState := make([]int, 10, 10)

		go func() {
			for {
				select {
				case value := <-stateChan:
					fmt.Println(waitGroup, ">>>>>>>>>>>>")
					newState[value.x] = value.l
					fmt.Println(value.l, value.x)
					fmt.Println(waitGroup, "<<<<<<<<<<<<")
					waitGroup.Done()
					fmt.Println(waitGroup, "============")
				}
			}
		}()

		waitGroup.Wait()
		fmt.Println(newState, "++++++++")
		previousState = newState
	}
}

func computeCell(nextState chan cell, currentState cell, neighbours ...int) {
	count := 0
	for _, value := range neighbours {
		if value == 1 {
			count++
		}
	}

	if count == 1 {
		nextState <- cell{x: currentState.x, l: 1}
	} else {
		nextState <- cell{x: currentState.x, l: 0}
	}
}
