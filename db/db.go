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

func Get(db *leveldb.DB, key []string) {
	var i= 0

	endSignal := make(chan bool, 1)
	go sleep(seconds, endSignal)

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
		default:
			db.Get([]byte(key[i]), nil)
			i++
		}
	}
}

func Put(db *leveldb.DB, key []string, value []string) {
	var i= 0

	endSignal := make(chan bool, 1)
	go sleep(seconds, endSignal)

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
		default:
			db.Put([]byte(key[i]), []byte(value[i]), nil)
			i++
		}
	}
}

func Delete(db *leveldb.DB, key []string) {
	var i = 0

	endSignal := make(chan bool, 1)
	go sleep(seconds, endSignal)

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
		default:
			db.Delete([]byte(key[i]), nil)
			i++
		}
	}
}

func sleep(seconds int, endSignal chan<- bool) {
	time.Sleep(time.Duration(seconds) * time.Second)
	endSignal <- true
}