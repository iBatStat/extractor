package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/iBatStat/extractor/db"
	myHttp "github.com/iBatStat/extractor/http"
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

	// start a http server and add all the relevant handlers

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yes monkey boy"))
	})
	http.Handle("/login", http.HandlerFunc(myHttp.LoginHandlerFunc))
	http.Handle("/newUser", http.HandlerFunc(myHttp.NewUserHandlerFunc))
	http.Handle("/uploadStat", myHttp.AuthenticateHandlerFunc(http.HandlerFunc(myHttp.UploadImageHandlerFunc)))

	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
	//	stat, err := san.ExtractFeatures("7splusBattery.jpeg")
	//	if err != nil {
	//		log.Fatal(err)
	//	} else {
	//		db.DBAccess.Push(stat)
	//		fmt.Println(fmt.Sprintf("****** Structured data is *********\n%s", *stat))
	//	}
}
