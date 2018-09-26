package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"../random"
	"strconv"
)

func CreateKeyValues(count int, minSize int, maxSize int)  {
	done1 := make(chan struct{})
	done2 := make(chan struct{})
	done3 := make(chan struct{})
	done4 := make(chan struct{})
	done5 := make(chan struct{})
	done6 := make(chan struct{})
	done7 := make(chan struct{})
	done8 := make(chan struct{})
	done9 := make(chan struct{})
	done10 := make(chan struct{})

	streams := count / 10;

	go func() {
		var str1 string
		var str2 string
		for i := 0; i < streams; i++ {
			str1, str2 = generate(str1, str2, minSize, maxSize, i, 1)
		}
		done1 <- struct{}{}
	}()

	go func() {
		var str1 string
		var str2 string
		for i := 0; i < streams; i++ {
			str1, str2 = generate(str1, str2, minSize, maxSize, i, 2)
		}
		done2 <- struct{}{}
	}()

	go func() {
		var str1 string
		var str2 string
		for i := 0; i < streams; i++ {
			str1, str2 = generate(str1, str2, minSize, maxSize, i, 3)
		}
		done3 <- struct{}{}
	}()

	go func() {
		var str1 string
		var str2 string
		for i := 0; i < streams; i++ {
			str1, str2 = generate(str1, str2, minSize, maxSize, i, 4)
		}
		done4 <- struct{}{}
	}()

	go func() {
		var str1 string
		var str2 string
		for i := 0; i < streams; i++ {
			str1, str2 = generate(str1, str2, minSize, maxSize, i, 5)
		}
		done5 <- struct{}{}
	}()

	go func() {
		var str1 string
		var str2 string
		for i := 0; i < streams; i++ {
			str1, str2 = generate(str1, str2, minSize, maxSize, i, 6)
		}
		done6 <- struct{}{}
	}()

	go func() {
		var str1 string
		var str2 string
		for i := 0; i < streams; i++ {
			str1, str2 = generate(str1, str2, minSize, maxSize, i, 7)
		}
		done7 <- struct{}{}
	}()

	go func() {
		var str1 string
		var str2 string
		for i := 0; i < streams; i++ {
			str1, str2 = generate(str1, str2, minSize, maxSize, i, 8)
		}
		done8 <- struct{}{}
	}()

	go func() {
		var str1 string
		var str2 string
		for i := 0; i < streams; i++ {
			str1, str2 = generate(str1, str2, minSize, maxSize, i, 9)
		}
		done9 <- struct{}{}
	}()

	go func() {
		var str1 string
		var str2 string
		for i := 0; i < streams; i++ {
			str1, str2 = generate(str1, str2, minSize, maxSize, i, 	10)
		}
		done10 <- struct{}{}
	}()

	<- done1
	<- done2
	<- done3
	<- done4
	<- done5
	<- done6
	<- done7
	<- done8
	<- done9
	<- done10
}

func WriteAllInOneFile(countOfFiles int) {
	for i := 1; i < countOfFiles; i++ {
		keys := Read("keys"+strconv.Itoa(i)+".txt")
		values := Read("values"+strconv.Itoa(i)+".txt")
		mainKeys := Read("keys.txt")
		mainValues := Read("values.txt")

		accKeys := mainKeys + "," + keys
		accValues := mainValues + "," + values

		Write("keys.txt", accKeys)
		Write("values.txt", accValues)

		println(i)
	}

}

func generate(str1 string, str2 string, minSize int, maxSize int, i int, streamNum int) (string, string)  {
	rndString1 := random.RandString(32)
	rndString2 := random.RandString(random.RandInt(minSize, maxSize))

	str1 += rndString1 +  ","
	str2 += rndString2 +  ","

	if i % 10000 == 0 {
		k := Read("keys"+strconv.Itoa(streamNum)+".txt")
		v := Read("values"+strconv.Itoa(streamNum)+".txt")

		nk := k + str1
		nv := v + str2

		str1 = ""
		str2 = ""

		Write("./keys"+strconv.Itoa(streamNum)+".txt", nk)
		Write("./values"+strconv.Itoa(streamNum)+".txt", nv)
		println(i*10)
	}

	return str1, str2
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