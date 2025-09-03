package main

import "fmt"

func main() {
	fmt.Println("We will be studying Pointers now");
	

	mynumber := 23;

	ptr := &mynumber;

	fmt.Println(ptr);
	fmt.Println(*ptr);

}