package main

import (
	"./db"
	"./file"
	"os"
	"strconv"
	"strings"
)

func main() {
	db.Connect()

	var seconds, err = strconv.Atoi(os.Getenv("SECONDS"))
	file.CreateKeyValues(seconds*10000000, 32, 100)

	if err != nil {
		panic(err.Error())
	}

	d1 := file.Read("keys.txt")
	d2 := file.Read("values.txt")

	keys := strings.Split(d1, ",")
	values := strings.Split(d2, ",")
	println("Count of keys: " + strconv.Itoa(len(keys)))
	println("Count of values: " + strconv.Itoa(len(values)))

	db.Put(db.Connect(), keys, values)

	db.Get(db.Connect(), keys)

	db.Delete(db.Connect(), keys)
}

