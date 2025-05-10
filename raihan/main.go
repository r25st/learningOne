package main

import (
	"fmt";
	"errors"
)

func main(){
	// var a string = "awuk ah";
	// printMe(a)
	a := 3
	b := 2

	// fmt.Println(addMe(a,b))
	// fmt.Println(subMe(a,b))
	result, reminder, err := divMe(a,b)
	// if(err!=nil){
	// 	fmt.Println(err.Error())
	// }else if(reminder == 0){
	// 	fmt.Printf("%d divided by %d is %d",a,b,result)
	// }else{
		// 	fmt.Printf("%d divided by %d is %d with a reminder of %d",a,b,result,reminder)
		
		// }
	switch{
		case err!=nil:
			fmt.Println(err.Error())
		case reminder == 0:
			fmt.Printf("%d divided by %d is %d",a,b,result)
		default:
			fmt.Printf("%d divided by %d is %d with a reminder of %d",a,b,result,reminder)
	}
}

func addMe(nemurator int, denominator int) int{
	return nemurator + denominator
}

func subMe(nemurator int, denominator int) int{
	return nemurator - denominator
	
}

func divMe(numurator int, denominator int) (int,int,error){
	var err error;
	if denominator == 0{
		err = errors.New("a number cannot divided by zero")
		return 0,0, err
	}
	var divisionResult int = numurator / denominator
	var divisionReminder int = numurator % denominator
	return divisionResult, divisionReminder, err
	
}



// func printMe(print string){
// 	fmt.Println(print)
// }