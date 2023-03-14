// To run this app we use --> go run HelloWorld.go
// But to print this statement we have to find our Print method which is present in fmt package so we will import fmt and using it we can print.
// As like in other language when we run go app then it first try to find out main method with -> "func"
// But we can have multiple main methods so how to find the entry point ? -> To find this entry point we will declare a package main and while
// running, it will first find main package and then inside it it will go for main method and except it as a entry point.

package main

import "fmt"

func main() {

	fmt.Print("Hello world")
}
