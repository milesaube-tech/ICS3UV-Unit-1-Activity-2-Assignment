/*
* Author: Miles Aube 
* Version: 1.0.0
* Date: 2025-10-02
* These are my programs to solve the mathematical equations
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	// Program 1: Convert 34°C to Fahrenheit
	fmt.Println("// The conversion of 34°C to Fahrenheit")
	fmt.Println("34°C in Fahrenheit = " + fmt.Sprint((9.0/5.0)*34 + 32))

	// Program 2: Area of a circle with radius 14 cm
	fmt.Println("\n// The area of a circle with radius 14 cm")
	fmt.Println("Area = " + fmt.Sprint(math.Pi*14*14) + " cm²")

	// Program 3: Volume of a sphere with radius 4 cm
	fmt.Println("\n// The volume of a sphere with radius 4 cm")
	fmt.Println("Volume = " + fmt.Sprint((4.0/3.0)*math.Pi*4*4*4) + " cm³")

	// Program 4: Gross pay for Fred (40 hours at $8.20/hour)
	fmt.Println("\n// Gross pay for Fred")
	fmt.Println("Fred's gross pay = " + fmt.Sprint(40*8.20) + " dollars")

	fmt.Println("\nDone.")
}