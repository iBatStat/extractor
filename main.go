package main

import (
	"fmt"
	"log"

	tes "github.com/otiai10/gosseract"
)

func main() {
	fmt.Println("starting new client")
	client, err := tes.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	var out string
	out, err = client.Src("/Users/adbhasin/Desktop/adi-iphone.jpeg").Out()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(out)
}
