package main

import (
	"example/hello/greetings"
	"example/hello/morestrings"

	"fmt"

	"github.com/google/go-cmp/cmp"
	"rsc.io/quote"
)



func main() {

	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())



	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	fmt.Println(greetings.Hello("sari"))

}


