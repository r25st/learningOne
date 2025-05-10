package calculation

import (
	"errors"
)

func Add(a, b int) int {
    return a + b
}
func Subtract(a,b int) int{
	return a - b
}
func Multiply(a,b int) int{
	return a * b
}
func Divide(a,b int) (int,int, error){
	if (a == 0 || b == 0){
		err := errors.New("cannot divided by zero")
		return 0, 0, err
	}
	result := a / b
	reminder := a % b
	return result, reminder, nil
}