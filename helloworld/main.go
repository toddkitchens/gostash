package main

import "fmt"

var y = 42
var a = `James said, "shaken
not
		stirred."`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
