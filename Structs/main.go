package main

import "fmt"

func main() {
	fmt.Println("We are goint to learn struct in go ");

	aayush := User{
		"Aayush", "Aayush@Verma.com", true, 21,
	}

	fmt.Println(aayush);

}


type User struct {
	Name string
	Email string
	Status bool
	age int
}


// Really Really simple it was