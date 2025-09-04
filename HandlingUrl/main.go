package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://karamjeetsony.tech";


func main() {
	fmt.Println("Let's Learn url handling in go");

	resultUrl, _ := url.Parse(myurl);

	fmt.Println(resultUrl);
	fmt.Println(resultUrl.ForceQuery)
	fmt.Println(resultUrl.Fragment)
	fmt.Println(resultUrl.Host)
	fmt.Println(resultUrl.OmitHost)
	fmt.Println(resultUrl.Opaque)
	fmt.Println(resultUrl.Port());
	fmt.Println(resultUrl.Path)
	fmt.Println(resultUrl.RawQuery)
	fmt.Println(resultUrl.RawPath)
	fmt.Println(resultUrl.Scheme)
	fmt.Println(resultUrl.User)
}