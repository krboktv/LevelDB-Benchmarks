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

	keys := strings.Split(d1, ",")
	values := strings.Split(d2, ",")
	println("Count of keys: " + strconv.Itoa(len(keys)))
	println("Count of values: " + strconv.Itoa(len(values)))

	db.Put(db.Connect(), keys, values)

	db.Get(db.Connect(), keys)

	db.Delete(db.Connect(), keys)
}

