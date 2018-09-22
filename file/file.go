package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"../random"
)

func CreateKeyValues(count int, minSize int, maxSize int)  {
	var str1 string
	var str2 string

	for i := 0; i < count; i++ {
		rndString1 := random.RandString(32)
		rndString2 := random.RandString(random.RandInt(minSize, maxSize))

		str1 += rndString1 +  ","
		str2 += rndString2 +  ","

		if i % 10000 == 0 {
			k := Read("keys.txt")
			v := Read("values.txt")

			nk := k + str1
			nv := v + str2

			str1 = ""
			str2 = ""

			Write("./keys.txt", nk)
			Write("./values.txt", nv)
			println(i)
		}
	}
}

func Write(fileName string, text string)  {
	file, err := os.Create(fileName)

	if err != nil{
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)
}

func Read(fileName string) string {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		os.Create(fileName)
	}

	return string(file)
}