package main

import (
	"errors"
	"fmt"
	strings2 "strings"
)

func concatStringsBad(string1, string2 string) (string, error) {
	var resString string

	if string1 == "" {
		return "", errors.New("string1 is empty")
	} else {
		if string2 == "" {
			return "", errors.New("string2 is empty")
		} else {
			resString = string1 + string2
		}
	}
	return resString, nil
}

func concatStringsGood(s1, s2 string, maxLength int) (string, error) {
	if s1 == "" {
		return "", errors.New("string1 is empty")
	}
	if s2 == "" {
		return "", errors.New("string2 is empty")
	}

	resString, err := concatString(s1, s2)
	if err != nil {
		return "", err
	}

	if len(resString) > maxLength {
		return resString[:maxLength], nil
	}

	return resString, nil

}

func concatString(s1, s2 string) (string, error) {
	if len(s1)+len(s2) < len(s1) || len(s1)+len(s2) < len(s2) {
		return "", errors.New("concatenation overflow")
	}
	return s1 + s2, nil
}

func join(strings ...string) (string, error) {
	if len(strings) == 0 {
		return "", errors.New("nothing to join, empty input")
	}

	var sb strings2.Builder

	for _, s := range strings {
		sb.WriteString(s)
	}
	return sb.String(), nil
}

func main() {
	s1 := "Hello"
	s2 := "WorldSomeAdditionalInformation"

	concatWord, err := concatStringsGood(s1, s2, 15)
	if err != nil {
		err.Error()
	}

	fmt.Println(concatWord)

	joinStrings, err := join(s1, s2)
	if err != nil {
		err.Error()
	}

	fmt.Println(joinStrings)
}
