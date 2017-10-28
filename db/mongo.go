package db

import (
	"crypto/tls"
	"net"

	"github.com/iBatStat/extractor/model"
	mgo "gopkg.in/mgo.v2"
)

type mongoAccess struct {
	c *mgo.Collection
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
	d.c = session.DB("istats").C("battery")
	return nil
}

func (d *mongoAccess) Push(batStat *model.BatteryStats) error {
	return d.c.Insert(batStat)
}

func (d *mongoAccess) getStatsForUser(user string) interface{} { return nil }
