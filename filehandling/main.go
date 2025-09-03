package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file handling in go lang");


	content := "This is going into some files okay..";

	file, err := os.Create("./gen.txt");

	if(err != nil) {
		panic(err);
	}


	length, writingError := io.WriteString(file, content);


	if(writingError != nil) {
		panic(writingError);
	}

	fmt.Println(length);

	defer file.Close();

	readFile("./gen.txt");

}



func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName);
	if(err != nil) {
		panic(err);
	}

	fmt.Println(string(data));


}