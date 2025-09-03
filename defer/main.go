package main

import "fmt"

func main() {
	defer fmt.Println("Hello From Aayush");


	// jab without defer wale sare execute ho jaye tab defer wala execute hoga.

	// defer last in first out strategy follow karta hai just like stack.


	// imagine just a stack every time any defer comes it pushed in to the stack.

	// har function ke complete code hone par uske saare push kiye gaye defer run and then popped from the stack.

	defer fmt.Println("and What is this?");

	fmt.Println("How are You??");

	fmt.Println("I knew it?");

	myDefer();

	just();

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i);
	}
}

func just() {
	defer fmt.Println("This is World's best line.");

	fmt.Println("Yes obviously");
}