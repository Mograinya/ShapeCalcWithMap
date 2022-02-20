package main

import (
	"fmt"
	"math"
	"strings"
)

type shape struct {
	dimensions         []float64
	inputPromptMessage string
	inputRequestPrompt func(sh shape)
	calcArea           func(sh shape) float64
	calcPerimeter      func(sh shape) float64
}

var knownShapesMap = make(map[string]shape)

func initShapeMap() {
	knownShapesMap["circle"] = shape{
		dimensions:         []float64{0},
		inputPromptMessage: "Input the radius of a circle:",
		inputRequestPrompt: func(sh shape) { fmt.Scanf("%f", &sh.dimensions[0]) },
		calcArea:           func(sh shape) float64 { return math.Pi * sh.dimensions[0] * sh.dimensions[0] },
		calcPerimeter:      func(sh shape) float64 { return 2 * math.Pi * sh.dimensions[0] },
	}

	knownShapesMap["square"] = shape{
		dimensions:         []float64{0},
		inputPromptMessage: "Input the side length of a square:",
		inputRequestPrompt: func(sh shape) { fmt.Scanf("%f", &sh.dimensions[0]) },
		calcArea:           func(sh shape) float64 { return sh.dimensions[0] * sh.dimensions[0] },
		calcPerimeter:      func(sh shape) float64 { return 4 * sh.dimensions[0] },
	}

	knownShapesMap["rectangle"] = shape{
		dimensions:         []float64{0, 0},
		inputPromptMessage: "Input length and width of a rectangle (separated by space):",
		inputRequestPrompt: func(sh shape) { fmt.Scanf("%f %f", &sh.dimensions[0], &sh.dimensions[1]) },
		calcArea:           func(sh shape) float64 { return sh.dimensions[0] * sh.dimensions[1] },
		calcPerimeter:      func(sh shape) float64 { return 2 * (sh.dimensions[0] + sh.dimensions[1]) },
	}

	knownShapesMap["triangle"] = shape{
		dimensions:         []float64{0, 0, 0},
		inputPromptMessage: "Input three sides of a triangle (separated by space):",
		inputRequestPrompt: func(sh shape) {
			fmt.Scanf("%f %f %f", &sh.dimensions[0], &sh.dimensions[1], &sh.dimensions[2])
		},
		calcArea: func(sh shape) float64 {
			d := sh.dimensions
			p := (d[0] + d[1] + d[2]) / 2
			return math.Sqrt(p * (p - d[0]) * (p - d[1]) * (p - d[2]))
		},
		calcPerimeter: func(sh shape) float64 {
			return sh.dimensions[0] + sh.dimensions[1] + sh.dimensions[2]
		},
	}
}

func requestShapeName() string {
	var key string
	fmt.Println("Enter the name of a shape to calculate:")
	fmt.Scanf("%s", &key)
	return strings.ToLower(key)
}

func printResult(sh shape) {
	fmt.Printf("\nCalculated area: %f", sh.calcArea(sh))
	fmt.Printf("\nCalculated perimeter: %f", sh.calcPerimeter(sh))
}

//----------------------------------------------------------------

func main() {
	initShapeMap()
	shapeName := requestShapeName()
	sh, ok := knownShapesMap[shapeName]
	if !ok {
		fmt.Println("Unknown shape")
		return
	}
	fmt.Println(sh.inputPromptMessage)
	fmt.Scanln() //Без этого следующий Scanf обнаруживает \n от прошлого ввода и выдаёт ошибку
	sh.inputRequestPrompt(sh)
	printResult(sh)
}
