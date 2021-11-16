package main // declare a main package

import (
	"fmt"
	"time"
) // import fmt package, contains functinos for formatting text

func main(){    // implement main function
	//fmt.Println(quote.Get("test"));
	fmt.Println(time.Now().GoString());
}