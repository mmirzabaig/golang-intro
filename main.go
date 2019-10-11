package main

import "fmt"

func main() {
	// #1
	fmt.Println(isOldEnoughToDrink(21))
	// #2
	objq2 := make(map[string]string)
	objq2["yes"] = "true"
	objq2["no"] = "false"
	fmt.Println(getProperty(objq2, "no"))
}

// #1
// Write a function called “isOldEnoughToDrink”.
// Given a number, in this case an age, “isOldEnoughToDrink” returns whether a person of this given age is old enough to legally drink in the United States.
// Notes:
// * The legal drinking age in the United States is 21.

func isOldEnoughToDrink(age int) bool {
	if age >= 21 {
		return true
	}
	return false
}

// Write a function called “getProperty”.
// Given an object and a key, “getProperty” returns the value of the property at the given key.
// Notes:
// * If there is no property at the given key, it should return undefined.

// var obj = {
//   key: 'value'
// };
// var output = getProperty(obj, 'key');
// console.log(output); // --> 'value'

func getProperty(obj map[string]string, str string) string {
	// if statements in go can include both a condition and an intialization
	// val will receive the value and ok will be either true or false
	if val, ok := obj[str]; ok {
		if ok == true {
			fmt.Println(val)
		} else {
			return "undefined"
		}
	}
	return ""
}
