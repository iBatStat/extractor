package main

import (
	"fmt"
	san "github.com/iBatStat/extractor/sanitizer"
	"log"
)

func main() {
	stat, err := san.ExtractFeatures("7splusBattery.jpeg")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(fmt.Sprintf("****** Extracted data is *********\n%v", *stat))
	}
}
