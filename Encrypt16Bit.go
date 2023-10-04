package main

import "fmt"

func convert2ASCII(input string) []uint16 {
	asciiRepresentation := ""

	for _, char := range input {
		asciiValue := int(char)
		temp := fmt.Sprintf("%b", asciiValue)

		for i := len(temp); i < 8; i++ {
			temp = "0" + temp
		}

		asciiRepresentation += temp
	}
	res := []uint16{}
	for i := range asciiRepresentation {
		res = append(res, uint16(asciiRepresentation[i]-'0'))
	}

	//fmt.Println("ASCII representation:", asciiRepresentation)
	return res
}

func encryptMsg(msg string, key []uint16) (res []uint16) {
	msgArr := convert2ASCII(msg)
	for i := range msgArr {
		res = append(res, msgArr[i]^key[i])
	}
	//fmt.Println("The encrypted msg is ", res)
	return res
}
