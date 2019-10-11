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
	// #3
	objq3 := make(map[string]bool)
	fmt.Println(addProperty(objq3, "rain"))
	// #4
	objq4 := make(map[string]int)
	objq4["odd"] = 3
	objq4["even"] = 4
	fmt.Println(removeProperty(objq4, "even"))
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

// #2
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

// #3
// Write a function called “addProperty”.
// Given an object, and a key, “addProperty” sets a new property on the given object with a value of true.

// var myObj = {};
// addProperty(myObj, 'myProperty');
// console.log(myObj.myProperty); // --> true

func addProperty(obj map[string]bool, key string) map[string]bool {
	obj[key] = true
	return obj
}

// #4
// Write a function called “removeProperty”.
// Given an object and a key, “removeProperty” removes the given key from the given object.

// var obj = {
//   name: 'Sam',
//   age: 20
// }
// removeProperty(obj, 'name');
// console.log(obj.name); // --> undefined

func removeProperty(obj map[string]int, even string) map[string]int {
	delete(obj, even)
	return obj
}
