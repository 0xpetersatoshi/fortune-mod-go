package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

const dataFileDirectory string = "../datfiles"

func main() {

	// reading the directory
	files, err := ioutil.ReadDir(dataFileDirectory)
	if err != nil {
		fmt.Println("Could not read dataFileDirectory:", err)
		return
	}

	numFiles := len(files)

	// create random seed
	randomSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randomSource)

	// check that there are files in directory
	if numFiles == 0 {
		fmt.Println("No files were present in the dataFileDirectory.")
		return
	}

	randomFileIndex := random.Intn(numFiles)

	// path to random file chosen
	fullPath := fmt.Sprintf("%s/%s", dataFileDirectory, files[randomFileIndex].Name())

	// read the random file
	bs, err := ioutil.ReadFile(fullPath)
	if err != nil {
		fmt.Println("Could not read file ", fullPath, ": ", err)
		return
	}

	// convert byte stream to string
	fortuneFileContents := string(bs)

	// split the fortunes in file by delimiter
	fortunes := strings.Split(fortuneFileContents, "%")
	numFortunes := len(fortunes)

	// check that there are fortunes in file
	if numFortunes == 0 {
		fmt.Println(fullPath, " did not contain any fortunes")
		return
	}

	// chose a random fortune
	randomFortuneIndex := random.Intn(numFortunes)
	fortune := fortunes[randomFortuneIndex]

	// print fortune to stdout
	fmt.Println(fortune)
}
