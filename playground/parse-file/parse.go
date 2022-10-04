package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type FieldLength struct {
	start uint64
	stop  uint64
}

type FieldData struct {
	lineNumber uint64
	data       string
}

const fields string = "ABCD,EFGH,IJKL"
const lengths string = "1-10,11-10,21-10"

func main() {
	fileNameArg, err := getFileName()
	checkForErr(err)

	lines, err := readFileLines(fileNameArg) 
	checkForErr(err)

	fieldLengths, err := getFieldLengthMap(fields, lengths, false)
	checkForErr(err)

	dataMap := getDataMap(lines, fieldLengths)
	fmt.Println(dataMap)
}

func getFileName() (string, error) {
	fileName := os.Args

	if len(fileName) < 2 {
		return "", errors.New("file name is required")
	}

	fmt.Println(fileName)
	return fileName[1], nil
}

func getDataMap(lines []string, fieldLengths map[string]FieldLength) ([]map[string]FieldData) {
	
	dataMap := []map[string]FieldData{}
	
	for i, line := range lines {
		theMap := make(map[string]FieldData)
		for key, value := range fieldLengths {

			if uint64(len(line)) < value.stop {
				value.stop = uint64(len(line))
			}

			theMap[key] = FieldData{uint64(i), line[value.start:value.stop]}
		}
		dataMap = append(dataMap, theMap)
	}
	return dataMap
}

func checkForErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readFileLines(fileName string) ([]string, error) {

	readFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	readFile.Close()
	return lines, err
}

func getFieldLengthMap(fields string, lengths string, zeroIndexed bool) (map[string]FieldLength, error) {

	fArr := strings.Split(fields, ",")
	lArr := strings.Split(lengths, ",")

	if len(fArr) != len(lArr) {
		return map[string]FieldLength{}, errors.New("jagged field length map")
	}

	fieldDataMap := map[string]FieldLength{}
	for i, str := range fArr {
		fieldLength, err := getFieldLength(lArr[i], zeroIndexed)
		if err != nil {
			return map[string]FieldLength{}, errors.New("invalid length map")
		}
		fieldDataMap[str] = fieldLength
	}

	return fieldDataMap, nil
}

func getFieldLength(fieldLength string, zeroIndex bool) (FieldLength, error) {
	lArr := strings.Split(fieldLength, "-")

	start, err := strconv.Atoi(lArr[0])
	if err != nil {
		log.Fatal(err)
		return FieldLength{}, err
	}

	if zeroIndex {
		start -= 1
	}

	length, err := strconv.Atoi(lArr[1])
	if err != nil {
		log.Fatal(err)
		return FieldLength{}, err
	}

	return FieldLength{start: uint64(start), stop: uint64(start + length)}, nil
}
