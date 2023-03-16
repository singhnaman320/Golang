package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Learning switch case with Golang..")

	// In Golang, the rand.Seed() function is used to set a seed value to generate pseudo-random numbers.
	// If the same seed value is used in every execution, then the same set of pseudo-random numbers is generated.
	// In order to get a different set of pseudo-random numbers, we need to update the seed value.

	// seed to get different result every time
	rand.Seed(time.Now().UnixNano())

	rollingDiceNumber := rand.Intn(6) + 1

	fmt.Println("Present number on dice is", rollingDiceNumber)

	switch rollingDiceNumber {

	case 1:
		fmt.Println("Move 1 spot..")

	case 2:
		fmt.Println("Move 2 spots..")

	case 3:
		fmt.Println("Move 3 spots..")

	case 4:
		fmt.Println("Move 4 spots..")
		fallthrough

	case 5:
		fmt.Println("Move 5 spots..")

	case 6:
		fmt.Println("Move 6 spots and take another chance..")

	default:
		fmt.Println("No face of dice can have more tan 6 points..")
	}

	// In Go, the control comes out of the switch statement immediately after a case is executed. A fallthrough statement is used to
	// transfer control to the first statement of the case that is present immediately after the case which has been executed.

	// The control comes inside the switch and the cases are evaluated. case 4: in line no. 35 is true and the program prints move 4 spots..
	// The next statement is fallthrough. When fallthrough is encountered the control moves to the first statement of the next case and also
	// prints move 5 spots... The output of the program is :

	//Move 4 spots..
	//Move 5 spots..
}
