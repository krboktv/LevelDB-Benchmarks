package main

import (
	"./db"
	"./file"
	"strconv"
	"strings"
)

func main() {
	db.Connect()

	//var env, err = strconv.Atoi(os.Getenv("SECONDS"))
	//file.CreateKeyValues(env, 32, 100)

	d1 := file.Read("keys.txt")
	d2 := file.Read("values.txt")
	println("Count of keys: " + strconv.Itoa(len(d1)))
	println("Count of values: " + strconv.Itoa(len(d2)))

	keys := strings.Split(d1, ",")
	values := strings.Split(d2, ",")

	db.Put(db.Connect(), keys, values)

	db.Get(db.Connect(), keys)

	db.Delete(db.Connect(), keys)
}

