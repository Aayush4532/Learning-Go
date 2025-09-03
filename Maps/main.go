package main

import "fmt"

func main() {
	fmt.Println("We are goiing to learn Maps here");

	languages := make(map[string]string);

	languages["Aayush"] = "Verma";

	languages["Karamjeet"] = "Sony";

	fmt.Println(languages);

	fmt.Printf("%T", languages);


	fmt.Println(languages["Aayush"]);


	delete(languages, "Karamjeet");

	fmt.Println(languages);






}