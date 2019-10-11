package main

import "fmt"

func main() {
	// #1 isOldEnoughToDrink
	fmt.Println(isOldEnoughToDrink(21))
	// #2 getProperty
	objq2 := make(map[string]string)
	objq2["yes"] = "true"
	objq2["no"] = "false"
	fmt.Println(getProperty(objq2, "no"))
	// #3 addProperty
	objq3 := make(map[string]bool)
	fmt.Println(addProperty(objq3, "rain"))
	// #4 removeProperty
	objq4 := make(map[string]int)
	objq4["odd"] = 3
	objq4["even"] = 4
	fmt.Println(removeProperty(objq4, "even"))
	// #5 getLengthOfTwoWords
	fmt.Println(getLengthOfTwoWords("one", "two"))
	// #6 addArrayProperty
	objq6 := make(map[string][2]string)
	arrq6 := [2]string{"hello", "world"}
	fmt.Println(addArrayProperty(objq6, "arr", arrq6))
	// #7 getLastElement
	arrq8 := []string{"12", "34", "56"}
	fmt.Println(getLastElement(arrq8))
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

// #5
// Write a function called “getLengthOfTwoWords”.
// Given 2 words, “getLengthOfTwoWords” returns the sum of their lengths.

func getLengthOfTwoWords(str1 string, str2 string) int {
	return len(str1) + len(str2)
}

// #6
// Write a function called “addArrayProperty”.
// Given an object, a key, and an array, “addArrayProperty” sets a new property on the object at the given key, with its value set to the given array.

// var myObj = {};
// var myArray = [1, 3];
// addArrayProperty(myObj, 'myProperty', myArray);
// console.log(myObj.myProperty); // --> [1, 3]

func addArrayProperty(obj map[string][2]string, key string, array [2]string) map[string][2]string {
	obj[key] = array
	return obj
}

// #7
// Write a function called “getLastElement”.

// Given an array, “getLastElement” returns the last element of the given array.

// Notes:
// * If the given array has a length of 0, it should return ‘undefined’.

// var output = getLastElement([1, 2, 3, 4]);
// console.log(output); // --> 4

func getLastElement(arr []string) string {
	return arr[len(arr)-1]
}
