package sanitizer

import (
	//"github.com/iBatStat/extractor/model"
	"fmt"
	gotes "github.com/otiai10/gosseract"
	"log"
)

var (
	ocr *gotes.Client
)

func testOCR(filePath string) {
	out, err := ocr.Src(filePath).Out()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(fmt.Sprintf("****** Extracted data is *********\n%s", out))
}
func init() {
	var err error
	ocr, err = gotes.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}
}
func ExtractFeatures(imgPath string) {
	testOCR(imgPath)
}
