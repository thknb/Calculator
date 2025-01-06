package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Қарапайым калькулятор")
	fmt.Println("-----------------------")

	// Бірінші санды енгізу
	fmt.Print("Бірінші санды енгізіңіз: ")
	fmt.Scanln(&num1)

	// Арифметикалық амал таңдау
	fmt.Print("Амал таңданыз: ")
	fmt.Scanln(&operator)

	// Екінші санды енгізу
	fmt.Print("Екінші санды енгізіңіз:")
	fmt.Scanln(&num2)

	// Есептеу
	switch operator {
	case "+":
		fmt.Printf("Result: %2f\n", num1+num2)
	case "-":
		fmt.Printf("Result: %2f\n", num1-num2)
	case "*":
		fmt.Printf("Result: %2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Result: %2f\n", num1/num2)
		} else {
			fmt.Println("Error! амал дұрыс таңдалмады!")
		}
	}
}
