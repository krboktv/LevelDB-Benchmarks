package main

import (
	"./db"
	"./random"
	"os"
	"strconv"
)

func main() {
	db.Connect()

	var env, err = strconv.Atoi(os.Getenv("SECONDS"))
	arraysLength := env * 1000000

	if err != nil {
		panic(err.Error())
	}

	keys := make([]string, arraysLength)
	values := make([]string, arraysLength)

	for i := 0; i < arraysLength; i++ {
		rndString := random.RandomString(32)
		rndString1 := random.RandomString(1000)
		keys[i] = rndString
		values[i] = rndString1
	}

	db.Put(db.Connect(), keys, values, arraysLength)

	db.Get(db.Connect(), keys, arraysLength)

	db.Delete(db.Connect(), keys, arraysLength)
}

