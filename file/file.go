package file

import (
	"../random"
	"fmt"
	"io/ioutil"
	"os"
)

func CreateKeyValues(count int, minSize int, maxSize int)  {
	keyChan := make(chan string)
	valueChan := make(chan string)

	doneKeys := make(chan bool, 1)
	doneValues := make(chan bool, 1)

	go generateKeys(count, keyChan, doneKeys)
	go readAndWriteKeys(keyChan)

	go generateValues(minSize, maxSize, count, valueChan, doneValues)
	go readAndWriteValues(valueChan)

	for {
		select {
			case <-doneKeys:
				close(keyChan)
			case <-doneValues:
				close(valueChan)
			default:

		}
	}
}

func generateKeys(countOfKeys int, c chan string, done chan bool) {
	var str string

	oneInteraction := 10000
	countOfInteractions := countOfKeys/oneInteraction


	for j := 0; j < countOfInteractions; j++ {
		for i := 0; i < oneInteraction; i++ {

			rndString := random.RandString(32)

			str += rndString +  ","
		}
		c <- str
		if j == (countOfInteractions - 1) {
			done <- true
		}
	}
}

func generateValues(minSize int, maxSize int, countOfKeys int, c chan string, done chan bool) {
	var str string

	oneInteraction := 10000
	countOfInteractions := countOfKeys/oneInteraction

	for j := 0; j < countOfInteractions; j++ {
		for i := 0; i < oneInteraction; i++ {
			rndString := random.RandString(random.RandInt(minSize, maxSize))

			str += rndString + ","
		}
		c <- str
		if j == (countOfInteractions - 1) {
			done <- true
		}
	}
}

func readAndWriteKeys(c chan string) {
	for {
		k := readKeys()
		newK := k + <-c
		writeKeys(newK)
	}
}

func readAndWriteValues(c chan string) {
	for {
		v := readValues()
		newV := v + <-c
		writeValues(newV)
	}
}

func readKeys() (string) {
	return Read("keys.txt")
}

func readValues() (string) {
	return Read("values.txt")
}

func writeKeys(text string) {
	Write("./keys.txt", text)
}

func writeValues(text string) {
	Write("./values.txt", text)
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