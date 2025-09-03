package main

import "fmt"

func main() {
	fmt.Println("We are goint to learn struct in go ");

	aayush := User{
		"Aayush", "Aayush@Verma.com", true, 21,
	}

	fmt.Println(aayush);

	aayush.getStatus();

}


type User struct {
	Name string
	Email string
	Status bool
	age int
}


func (u User) getStatus() {
	fmt.Println(u.Status);
}