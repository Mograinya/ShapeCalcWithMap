package main

import (
	"fmt"
	"math"
	"strings"
)

type figureStruct struct {
	dimensions         []float64
	inputRequestPrompt func(fig figureStruct)
	calcArea           func(fig figureStruct) float64
	calcPerimeter      func(fig figureStruct) float64
}

var knownFiguresMap = make(map[string]figureStruct)

func initFigureMap() {
	knownFiguresMap["circle"] = figureStruct{
		dimensions: []float64{0},
		inputRequestPrompt: func(fig figureStruct) {
			fmt.Println("Input the radius of a circle:")
			fmt.Scanf("%f", &fig.dimensions[0])
		},
		calcArea:      func(fig figureStruct) float64 { return math.Pi * fig.dimensions[0] * fig.dimensions[0] },
		calcPerimeter: func(fig figureStruct) float64 { return 2 * math.Pi * fig.dimensions[0] },
	}

	knownFiguresMap["square"] = figureStruct{
		dimensions: []float64{0},
		inputRequestPrompt: func(fig figureStruct) {
			fmt.Println("Input the side length of a square:")
			fmt.Scanf("%f", &fig.dimensions[0])
		},
		calcArea:      func(fig figureStruct) float64 { return fig.dimensions[0] * fig.dimensions[0] },
		calcPerimeter: func(fig figureStruct) float64 { return 4 * fig.dimensions[0] },
	}

	knownFiguresMap["rectangle"] = figureStruct{
		dimensions: []float64{0, 0},
		inputRequestPrompt: func(fig figureStruct) {
			fmt.Println("Input length and width of a rectangle separated by space:")
			fmt.Scanf("%f %f", &fig.dimensions[0], &fig.dimensions[1])
		},
		calcArea:      func(fig figureStruct) float64 { return fig.dimensions[0] * fig.dimensions[1] },
		calcPerimeter: func(fig figureStruct) float64 { return 2 * (fig.dimensions[0] + fig.dimensions[1]) },
	}

	knownFiguresMap["triangle"] = figureStruct{
		dimensions: []float64{0, 0, 0},
		inputRequestPrompt: func(fig figureStruct) {
			fmt.Println("Input three sides of a tiangle separated by space:")
			fmt.Scanf("%f %f %f", &fig.dimensions[0], &fig.dimensions[1], &fig.dimensions[2])
		},
		calcArea: func(fig figureStruct) float64 {
			d := fig.dimensions
			p := (d[0] + d[1] + d[2]) / 2
			return math.Sqrt(p * (p - d[0]) * (p - d[1]) * (p - d[2]))
		},
		calcPerimeter: func(fig figureStruct) float64 {
			return fig.dimensions[0] + fig.dimensions[1] + fig.dimensions[2]
		},
	}
}

func figureTypeRequest() string {
	var key string
	fmt.Println("Enter type of figure:")
	fmt.Scanf("%s", &key)
	return strings.ToLower(key)
}

func printResult(fig figureStruct) {
	fmt.Printf("\nThe area of this figure is %f", fig.calcArea(fig))
	fmt.Printf("\nThe perimeter of this figure is %f", fig.calcPerimeter(fig))
}

//----------------------------------------------------------------

func main() {
	initFigureMap()
	figureType := figureTypeRequest()
	fig, ok := knownFiguresMap[figureType]
	if !ok {
		fmt.Println("Unknown figure")
		return
	}
	fig.inputRequestPrompt(fig)
	printResult(fig)
}
