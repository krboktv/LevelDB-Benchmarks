package db

import (
	"github.com/syndtr/goleveldb/leveldb"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

var seconds, err = strconv.Atoi(os.Getenv("SECONDS"))
var nanoseconds = seconds * 1e9

func Connect() (*leveldb.DB)  {
	db, err := leveldb.OpenFile("../db", nil)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	return db
}

func Get(db *leveldb.DB, key []string, count int) {
	var i = 0
	timer1 := time.NewTimer(time.Duration(nanoseconds))
	for ; i < count; i++ {
		db.Get([]byte(key[i]), nil)
	}
	<-timer1.C
	println("Get query (query per second)")
	qps := i / seconds
	println(qps)
	err := ioutil.WriteFile("get.txt", []byte(strconv.Itoa(qps)), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func Put(db *leveldb.DB, key []string, value []string, count int) {
	var i = 0
	timer1 := time.NewTimer(time.Duration(nanoseconds))
	for ; i < count; i++ {
		db.Put([]byte(key[i]), []byte(value[i]), nil)
	}
	<-timer1.C
	println("Put query (query per second)")
	qps := i / seconds
	println(qps)
	err := ioutil.WriteFile("put.txt", []byte(strconv.Itoa(qps)), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func Delete(db *leveldb.DB, key []string, count int) {
	var i = 0
	timer1 := time.NewTimer(time.Duration(nanoseconds))
	for ; i < count; i++ {
		db.Delete([]byte(key[i]), nil)
	}
	<-timer1.C
	println("Delete query (query per second)")
	qps := i / seconds
	println(qps)
	err := ioutil.WriteFile("delete.txt", []byte(strconv.Itoa(qps)), 0666)
	if err != nil {
		log.Fatal(err)
	}
}