package main

import (
	"./db"
	"./file"
	"strings"
)

func main() {
	db.Connect()

	//var env, err = strconv.Atoi(os.Getenv("SECONDS"))
 	//arraysLength := env * 1000000
	arraysLength := 2000000000
	//file.CreateKeyValues(arraysLength, 32, 100)

	d1 := file.Read("keys.txt")
	d2 := file.Read("values.txt")

	keys := strings.Split(d1, ",")
	values := strings.Split(d2, ",")

	db.Put(db.Connect(), keys, values, arraysLength)

	db.Get(db.Connect(), keys, arraysLength)

	db.Delete(db.Connect(), keys, arraysLength)
}

