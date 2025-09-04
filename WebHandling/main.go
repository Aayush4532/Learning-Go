package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://karamjeetsony.tech";

func main() {
	fmt.Println("Learning to create a web request");

	res, err := http.Get(url);

	if(err != nil) {
		panic(err);
	}

	databytes, err := io.ReadAll(res.Body);

	if(err != nil) {
		panic(err);
	}

	content := string(databytes);

	fmt.Println(content);


	defer res.Body.Close();


}