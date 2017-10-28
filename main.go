package main

import (
	"fmt"
	san "github.com/iBatStat/extractor/sanitizer"
	"log"
)

func main() {
	stat, err := san.ExtractFeatures("6sBattery.jpg")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(fmt.Sprintf("****** Extracted data is *********\n%s", *stat))
	}
}
