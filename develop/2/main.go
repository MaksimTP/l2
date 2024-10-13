package main

import (
	"fmt"
	"log"
)

func unpackString(s string) (res string, err error) {

	if len(s) == 0 {
		return "", nil
	}
	s, err = validateString(s)

	if err != nil {
		return "", err
	}

	rs := []rune(s)

	for i := 0; i < len(rs)-1; i += 2 {
		for j := 0; j < int(rs[i+1]-'0'); j++ {
			res += string(rs[i])
		}
	}

	return
}

func validateString(s string) (string, error) {
	rs := []rune(s)
	res := ""
	var isCurRune bool
	var isNextDigit bool
	for i := 0; i < len(rs)-1; i++ {
		isCurRune = rs[i] > '9' || rs[i] < '0'
		isNextDigit = rs[i+1] <= '9' && rs[i+1] >= '0'
		if !isCurRune && isNextDigit {
			return "", fmt.Errorf("incorrect string")
		}
		if rs[i] == 92 {
			i++
			isCurRune = true
			if i == len(rs)-1 {
				break
			}
			isNextDigit = rs[i+1] <= '9' && rs[i+1] >= '0'
		}

		if isCurRune && isNextDigit {
			res += string(rs[i])
			res += string(rs[i+1])
		} else if isCurRune && !isNextDigit {
			res += string(rs[i])
			res += "1"
		}
		isCurRune = false
		isNextDigit = false
	}

	if rs[len(rs)-1] > '9' || rs[len(rs)-1] < '0' || rs[len(rs)-1] == 92 || isCurRune {
		res += string(rs[len(rs)-1])
		res += "1"
	}

	return res, nil
}

func main() {
	s, err := unpackString(``)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(s)
}