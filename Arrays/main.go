package main

import "fmt"

func main() {
	fmt.Println("I am going to learn Arrays in go");

	var fruitlist [4]string;

	fruitlist[0] = "Apple";
	fruitlist[1] = "Mango";
	fruitlist[3] = "Banana";

	fmt.Println(fruitlist);
	fmt.Println(len(fruitlist));

	var veg = [3] string {"Potato", "Tomato", "Onion"};

	fmt.Println(veg);



}