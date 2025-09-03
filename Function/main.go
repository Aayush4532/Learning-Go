package main

import "fmt"

func main() {
	greeter();
	fmt.Println("Now I am Here in main function");
	result := sum(3, 4);
	fmt.Println(result);
	fmt.Println(proSum(2,4,4,5,5,1,332));
}

func greeter() {
	fmt.Println("Namaste From Golang");
}

func sum(a int, b int) int{
	return a + b;
}

func proSum(a ...int) int{
	total := 0;
	for i := 0; i < len(a); i++ {
		total += a[i];
	}

	return total;
}