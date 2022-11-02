package FileReadWork

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	checkHashstandard   string = "c00f0c4675b91fb8b918e4079a0b1bac"
	checkHashshadow     string = "3bf1114a986ba87ed28fc1b5884fc2f8"
	checkHashthinkertoy string = "b7dcd117ca3e521a1e443f05634b64f3"
)

const CountLinesInTxtFile = 760

func FileReadWork(fileName string, checkFileName string) ([]string, error) {
	if !CheckH(GetMD5Hash((CheckFormat(checkFileName))), checkHashshadow, checkHashstandard, checkHashthinkertoy) {
		log.Fatalln("Unacceptable .txt file")
		return []string{}, errors.New("FILE IS MODIFICATED HASH ERROR")
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error when read file")
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	text := ""
	counter, counterForappendSlice := 0, 0
	var massivString []string
	for scanner.Scan() {

		// fmt.Println(len(scanner.Text()))
		if len(scanner.Text()) == 0 {
			continue
		}
		counter++
		counterForappendSlice++
		text += string(scanner.Text())
		if counter != CountLinesInTxtFile {
			text += "\n"
		}
		if counterForappendSlice%8 == 0 {
			massivString = append(massivString, text)
			text = ""
		}

		// fmt.Println(scanner.Text())

	}
	// fmt.Println(massivString)

	if err := scanner.Err(); err != nil {
		fmt.Println("scanner error")
		return massivString, err

	}

	return massivString, nil
}

func CheckH(a, b, c, d string) bool {
	if a != b && a != c && a != d {
		return false
	}
	return true
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func CheckFormat(s string) string {
	if s == "shadow" {
		return "shadow"
	}
	if s == "standard" {
		return "standard"
	}
	if s == "thinkertoy" {
		return "thinkertoy"
	}
	return ""
}
