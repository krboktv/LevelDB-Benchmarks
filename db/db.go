package db

import (
	"../file"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
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
	var i= 0

	endSignal := make(chan bool, 1)
	go sleep(seconds, endSignal)

	pollInt := time.Second

	go func() {
		for ; i < count; i++ {
			db.Get([]byte(key[i]), nil)
		}
	}()

	for {
		select {
		case <-endSignal:
			println("Get query (query per second)")
			qps := i / seconds
			println(qps)
			file.Write("get.txt", strconv.Itoa(qps))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("The end!")
			return
		}
		time.Sleep(pollInt)
	}
}

func Put(db *leveldb.DB, key []string, value []string, count int) {
	var i= 0

	endSignal := make(chan bool, 1)
	go sleep(seconds, endSignal)

	pollInt := time.Second

	go func() {
		for ; i < count; i++ {
			db.Put([]byte(key[i]), []byte(value[i]), nil)
		}
	}()

	for {
		select {
		case <-endSignal:
			println("Put query (query per second)")
			qps := i / seconds
			println(qps)
			file.Write("put.txt", strconv.Itoa(qps))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("The end!")
			return
		}
		time.Sleep(pollInt)
	}
}

func Delete(db *leveldb.DB, key []string, count int) {
	var i= 0

	endSignal := make(chan bool, 1)
	go sleep(seconds, endSignal)

	pollInt := time.Microsecond

	go func() {
		for ; i < count; i++ {
			db.Delete([]byte(key[i]), nil)
		}
	}()

	for {
		select {
		case <-endSignal:
			println("Delete query (query per second)")
			qps := i / seconds
			println(qps)
			file.Write("delete.txt", strconv.Itoa(qps))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("The end!")
			return
		}
		time.Sleep(pollInt)
	}
}

func sleep(seconds int, endSignal chan<- bool) {
	time.Sleep(time.Duration(seconds) * time.Second)
	endSignal <- true
}