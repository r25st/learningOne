package main

import (
    "fmt"
    "learningOne/calculator/calculation"
)

func main() {

		var ques1 string
		var numerator int
		var denumerator int
	fmt.Println("~~ Welcome to the calculator ~~")
	fmt.Println("Which operation do you want to use?\na) Addition\nb) Subtraction\nc) Multiplication\nd) Divition")
    fmt.Scanln(&ques1)
	fmt.Println("Enter the numerator:")
	fmt.Scanln(&numerator)
	fmt.Println("Enter the denumerator:")
	fmt.Scanln(&denumerator)

	switch ques1{
	case "a", "A":
		result := calculation.Add(numerator,denumerator)
		fmt.Printf("The sum of %d and %d is %d", numerator,denumerator,result)
	
	case "b", "B":
		result := calculation.Subtract(numerator,denumerator)
		fmt.Printf("%d subtracted by %d is %d", numerator, denumerator, result)
		
	case "c", "C":
		result := calculation.Multiply(numerator,denumerator)
		fmt.Printf("%d Multiply by %d is %d", numerator, denumerator, result)
	
	case "d", "D":
		result,reminder,err := calculation.Divide(numerator,denumerator)
		if err != nil {
			fmt.Println("Error:", err.Error())
		} else if reminder == 0 {
			fmt.Printf("%d divided by %d is %d", numerator, denumerator, result)
		} else {
			fmt.Printf("%d divided by %d is %d with a remainder of %d", numerator, denumerator, result, reminder)
		}
		default:
			fmt.Println("Invalid, please try again")
		}
	}
