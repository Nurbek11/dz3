package dz3

import (
	"fmt"
	"github.com/jakehl/goid"
	"math"
	"unicode"
)

//package function
func CaseChanger(word string) string {
	newWord := ""
	//iterate the word symbols
	for _, r := range word {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			//concatenate the symbols
			newWord += string(unicode.ToLower(r))
		} else {
			//concatenate the symbols
			newWord += string(unicode.ToUpper(r))
		}
	}
	return newWord
}

func QuadraticRoot(a, b, c float64) {
	D := (b * b) - (4 * a * c)

	if D > 0 {
		root1 := -b + math.Sqrt(D)/(2*a)
		root2 := -b - math.Sqrt(D)/(2*a)
		fmt.Println("root1 = ", root1, ", root2 = ", root2)
	} else if D == 0 {
		root1 := -b / (2 * a)
		root2 := -b / (2 * a)
		fmt.Println("root1 = ", root1, ", root2 = ", root2)
	} else if D < 0 {
		root1 := -b / (2 * a)
		root2 := -b / (2 * a)
		Img := math.Sqrt(-D) / (2 * a)
		fmt.Println("root1 = ", root1, "+", Img, ", root2 = ", root2, "-", Img)
	}
}

func Uuid() {
	v4UUID := goid.NewV4UUID()
	fmt.Println(v4UUID)
}
