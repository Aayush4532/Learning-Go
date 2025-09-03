package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input"
	fmt.Println(welcome);

	reader := bufio.NewReader(os.Stdin);
	fmt.Print("Enter Your Name : ");

	input, _ := reader.ReadString('\n');
	fmt.Println("Sabka Papa -> : ", input);

	fmt.Printf("The type of input is : %T", input);

	fmt.Print("Go Corona Go");

}