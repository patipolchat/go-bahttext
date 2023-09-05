package main

import (
	"fmt"

	convert "github.com/patipolchat/go-bahttext/convert"
)

func main() {
	exampleAmount := "7,532.73 บาท"
	result, err := convert.ConvertToText(exampleAmount)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s => %s\n", exampleAmount, result)

}
