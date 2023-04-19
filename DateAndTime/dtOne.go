package main

import (
	"fmt"
	"time"
)

// In Go language, time packages supplies functionality for determining as well as viewing time. The Date() function in Go language is used to
// find Date the Time which is equivalent to yyyy-mm-dd hh:mm:ss + nsec nanoseconds in the proper time zone in the stated location. Moreover,
// this function is defined under the time package. Here, you need to import the “time” package in order to use these functions.

// Syntax:

// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time

// Here, “loc” is pointed to the location.

// Return Value: It returns a time that is proper in one of the two associated zones in the conversion, but it won’t guarantee that which one
// is returned. And it returns panics if the given “loc” is nil.

// Note: Here, the value of the month, day hour, min, sec and nsec can be more than the normal ranges but they will be automatically normalized
// during the conversion. For example, if the date is April 34 then it is converted to May 1.

func main() {

	givenTime := time.Date(2022, time.April, 10, 24, 60, 57, 00, time.UTC)

	fmt.Printf("Time is %s", givenTime.Local())
}
