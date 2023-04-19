package main

import (
	"fmt"
	"time"
)

func main(){

	// Calling Date() method with all its parameters
    tm := time.Date(2020, time.April, 34, 25, 72, 01, 0, time.UTC)

  	// Using Local() for location and printing the stated time and date in UTC
  	fmt.Printf("Date is %s", tm.Local())
}