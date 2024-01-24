package main

import (
	"fmt"
	"reflect"
)

func main() {

	var stringA = "This is a test!"
	var stringB = "is This test! A"

	if anagramCheck(stringA, stringB) {
		fmt.Printf("Strings are anagrams!\n")
	} else {
		fmt.Printf("Strings are unfortunately not anagrams!\n")
	}

	stringA = "This is a test!"
	stringB = "This is test! a"

	if anagramCheck(stringA, stringB) {
		fmt.Printf("Strings are anagrams!\n")
	} else {
		fmt.Printf("Strings are unfortunately not anagrams!\n")
	}

	return
}

// Do all the anagram checking things!
func anagramCheck(stringA string, stringB string) bool {

	// First check, make sure string length is equal.
	lenA := checkStringLength(stringA)
	lenB := checkStringLength(stringB)

	// If not equal, then return false
	if lenA != lenB {
		fmt.Printf("[%s] length %d and [%s] length %d are not anagrams.\n", stringA, lenA, stringB, lenB)
		return false
	}

	// Second, lets go through each string and count the chracters.
	stringMapA := buildMap(stringA)
	stringMapB := buildMap(stringB)

	// Verbose output of strings and length
	fmt.Printf("String A:[%s] length:%d  String B:[%s] length:%d\n", stringA, lenA, stringB, lenB)

	// Third, compare character count.
	eq := compareCharacterCount(stringMapA, stringMapB)

	if eq {
		return true
	}

	return false
}

// Are the dictionaries equal?
func compareCharacterCount(stringA map[rune]int, stringB map[rune]int) bool {
	return reflect.DeepEqual(stringA, stringB)
}

// Make a dictionary of each character and number of times found.
func buildMap(str string) map[rune]int {
	var stringMap = make(map[rune]int)
	for _, byteChar := range str {
		if _, ok := stringMap[byteChar]; !ok {
			stringMap[byteChar] = 1
		}
		stringMap[byteChar]++
	}
	return stringMap
}

// Return length of an input string.
func checkStringLength(str string) int {
	return len(str)
}
