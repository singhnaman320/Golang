package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learn date and time with Golang")
	parsentTime := time.Now()

	fmt.Println("\nTime one: ",parsentTime)
	fmt.Println("\nTime two: ",parsentTime.Format("01-02-2006 15:04:05 Monday"))
	cratedDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println("\nDate one",cratedDate)
	fmt.Println("\nDate Two",cratedDate.Format("01-02-2016 15:04:05 Monday"))
}