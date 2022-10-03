package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

type FieldLength struct {
	start uint64
	stop  uint64
}

func main() {

}

func GetFieldLengthMap(fields string, lengths string, zeroIndexed bool) (map[string]FieldLength, error) {

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
