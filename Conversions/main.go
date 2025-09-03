package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please Rate Our Pizza, Nahi to I will be Your Jija");
	fmt.Print("Rate b/w 1 and 5 : ");

	reader := bufio.NewReader(os.Stdin);

	input, _ := reader.ReadString('\n');

	fmt.Println("Input dekh ye aaya tha ", input);
	fmt.Printf("T ype dekh ye  tha : %T", input);

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64);
	fmt.Println();

	if(err != nil) {
		fmt.Println(err);
	} else {
		numRating++;
		fmt.Println("numRating is : " , numRating);
	}
}