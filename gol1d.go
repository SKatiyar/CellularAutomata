package main

import "fmt"
import "image"
import "os"
import "math"
import "image/jpeg"
import "image/color"

const size int = 1000
const numSeeds int = 21

func main() {
	birth := make([]int, size, size)
	emptyImage := image.NewRGBA(image.Rect(0, 0, size, size))

	imgFile, imgFileErr := os.Create("gol.jpg")
	if imgFileErr != nil {
		fmt.Println(imgFileErr)
		return
	}

	for i := 0; i < numSeeds; i++ {
		cell := int(math.Floor(math.Abs(math.Sin(float64(i))) * float64(size)))
		if cell != 0 && cell != size-1 {
			fmt.Println(cell)
			birth[cell] = 1
		}
	}

	for z := 0; z < size; z++ {
		result := make([]int, size, size)

		for i := 1; i < size-1; i++ {
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

		for i := 0; i < size; i++ {
			if birth[i] == 1 {
				emptyImage.Set(i, z, color.RGBA{R: 39, G: 174, B: 96, A: 0})
			} else {
				emptyImage.Set(i, z, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}

		birth = result
	}

	if err := jpeg.Encode(imgFile, emptyImage, &jpeg.Options{Quality: 100}); err != nil {
		fmt.Println(err)
		return
	}
}
