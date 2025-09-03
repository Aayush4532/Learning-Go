package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("We are going to study Slices in GO");

	var fruitList = [] string{};

	fmt.Printf("Type of fruitlist is : %T", fruitList);
	fmt.Println();

	fruitList = append(fruitList, "Mango", "Banana");

	fmt.Println(fruitList);


	highScore := make([]int, 4);

	highScore[0] = 938424;
	highScore[1] = 29;
	highScore[2] = 323844;
	highScore[3] = 55;

	highScore = append(highScore, 2873, 983, 2834, 9283); // this function does expand memory allocation


	fmt.Println(highScore);

	sort.Ints(highScore);

	fmt.Println(highScore);


	// removing elements with the help of indexes.

	var index int = 3;

	highScore = append(highScore[index:]);

	fmt.Println(highScore);





}