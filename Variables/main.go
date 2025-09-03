package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	c := 0
	c = a
	a = b
	b = c

	fmt.Println(a, b);

	var username string = "Aayush"
	fmt.Println(username);

	var isLegend bool = false;
	if(username == "Aayush") {
		isLegend = true;
	}


	if(isLegend) {
		fmt.Println("Prabhu aap hi legend ho");
	}
	

	var first float32 = 8293.0239
	var second float64 = 2.34

	fmt.Println(first / float32(second));

	k := 8.909

	fmt.Println(k);

	var noob = "chhodo"
	fmt.Println(noob);

	fmt.Println(noob);

}