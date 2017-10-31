package db

import (
	"crypto/tls"
	"log"
	"net"

	"github.com/iBatStat/extractor/model"
	mgo "gopkg.in/mgo.v2"
)

type mongoAccess struct {
	data  *mgo.Collection
	users *mgo.Collection
}

var DBAccess = new(mongoAccess)

func (d *mongoAccess) Init(user, password string, hosts []string) error {

	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	dialInfo := &mgo.DialInfo{
		Addrs:    hosts,
		Database: "admin",
		Username: user,
		Password: password,
	}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		return err
	}
	d.data = session.DB("istats").C("battery")
	d.users = session.DB("istats").C("users")
	return nil
}

func (d *mongoAccess) Push(batStat *model.BatteryStats) error {
	return d.data.Insert(batStat)
}

func (d *mongoAccess) GetUser(userEmail string) *model.User {
	var user model.User
	err := d.users.Find(map[string]interface{}{"email": userEmail}).One(&user)
	if err != nil {
		log.Println(err)
		return nil
	} else {
		return &user
	}

}

func (d *mongoAccess) SaveUser(user model.User) error {
	return d.users.Insert(user)

}
