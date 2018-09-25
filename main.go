package main

import (
	"./db"
	"fmt"
	"os"
	"strconv"
	"strings"
	"./file"
)

func main() {
	db.Connect()

	fmt.Println(os.Getenv("SECONDS"))
	var env, err = strconv.Atoi(os.Getenv("SECONDS"))
	file.CreateKeyValues(env * 10000000, 32, 100)

	if err != nil {
		panic(err.Error())
	}

	file.WriteAllInOneFile(20)

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

