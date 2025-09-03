package main

import (
	"fmt"
	"time"
)

func main() {
	presenttime := time.Now();

	fmt.Println(presenttime);

	fmt.Println(presenttime.Format("02-01-2006"));

	time := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC);

	fmt.Println(time);

}