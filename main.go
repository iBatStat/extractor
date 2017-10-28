package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/iBatStat/extractor/db"
	san "github.com/iBatStat/extractor/sanitizer"
)

func main() {
	var (
		user, password, hostsCombined string
	)
	flag.StringVar(&user, "dbUser", "", "Mongo db user")
	flag.StringVar(&password, "dbPass", "", "Mongo db pass")
	flag.StringVar(&hostsCombined, "dbHosts", "", "Comma seperated list of mongo db hosts. No spaces!!!")

	flag.Parse()
	hosts := strings.Split(hostsCombined, ",")
	err := db.DBAccess.Init(user, password, hosts)
	if err != nil {
		log.Fatal(err)
		return
	}
	stat, err := san.ExtractFeatures("7splusBattery.jpeg")
	if err != nil {
		log.Fatal(err)
	} else {
		db.DBAccess.Push(stat)
		fmt.Println(fmt.Sprintf("****** Structured data is *********\n%s", *stat))
	}
}
