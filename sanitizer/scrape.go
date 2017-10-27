package sanitizer

import (
	"fmt"
	"github.com/iBatStat/extractor/model"
	gotes "github.com/otiai10/gosseract"
	"log"
)

var (
	ocr *gotes.Client
)

func init() {
	var err error
	ocr, err = gotes.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}
}
func ExtractFeatures(imgPath string) (*model.BatteryStats, error) {
	out, err := ocr.Src(imgPath).Out()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(fmt.Sprintf("****** Extracted data is *********\n%s", out))
	//TO DO Use ocr to extract data nd the construct the BatterStats out of it
	return &model.BatteryStats{}, nil
}
