package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	input := bufio.NewReader(os.Stdin);
	num, _ := input.ReadString('\n');
	num = strings.TrimSpace(num);
	n, _ := strconv.Atoi(num);


	t := 2 * n;
	for i := 0; i < t; i++ {
		var temp int = i;
		if(i == n) {
			continue;
		}
		if(temp > n) {
			temp = t - i - 1;
		}
		for j := 0; j < t; j++ {
			if j <= temp || t - temp - 1 <= j {
				fmt.Print("*");
			} else {
				fmt.Print(" ");
			}
		}
		fmt.Println();
	}
}