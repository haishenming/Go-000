package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	err := GetUser()
	if IsRecordNotFind(err) {
		fmt.Println("do some thing for record not found")
	} else {
		fmt.Println("do some thing for other err")
	}
}

