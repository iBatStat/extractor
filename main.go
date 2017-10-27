package main

import (
	san "github.com/iBatStat/extractor/sanitizer"
)

func main() {
	san.ExtractFeatures("7splusBattery.jpeg")
	san.ExtractFeatures("6sBattery.jpg")
}
