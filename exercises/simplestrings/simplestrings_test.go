package simplestrings

import (
	"testing"
	"fmt"
)

const weekdays = "Monday Tuesday Wednesday Thursday Friday"

// test that Tuesday is a weekday
func TestTuesday(t *testing.T) {
	testweekday := "Tuesday"

 if Contains(weekdays,testweekday) {
 	fmt.Println(testweekday+" is on weekday")
 }

}
// test that Sunday is not a weekday
func TestSunday(t *testing.T) {
	testweekday := "Sunday"

	if !Contains(weekdays,testweekday) {
		fmt.Println(testweekday+" out of weekday")
	}

}


// test that an empty search string returns 0
func TestEmptySearch(t *testing.T) {
	testweekday := ""

	if Index(weekdays,testweekday)==0 {
		fmt.Println(" Return Zero")
	}

}

// test that the string Monday is not found in the empty string
func TestContains(t *testing.T) {
	testString := ""
	Monday := "Monday"
	if !Contains(testString,Monday){
		fmt.Println(" Return False")
	}
}