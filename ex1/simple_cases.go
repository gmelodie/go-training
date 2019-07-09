package ex1

import (
	"errors"
	"strconv"
	"strings"
)

//HelloWorld classic function
func HelloWorld() string {
	return "Hello World"
}

//AddInts return sum of two int
func AddInts(value1, value2 int) int {
	return value1 + value2
}

//AddIntsDifferents return sum of two differents ints (64 and 32 bits)
func AddIntsDifferents(value1 int64, value2 int32) int {
	return int(value1 + int64(value2))
}

//Substring standard method, don't use library
func Substring(value string, left, right int) string {
	if left > right {
		return ""
	}

	return value[left:right]
}

// SubstringWithErrors return substringed value but can generate errors
func SubstringWithErrors(value string, left, right int) (string, error) {

	if left < 0 || left >= len(value) || right < 0 || right >= len(value) || right < left {
		return "", errors.New("")
	}

	return value[left:right], nil
}

// ExtractNumbersFromString return the numbers existing in the list
func ExtractNumbersFromString(value string) []int {
	numbers := make([]int, 0)

	words := strings.FieldsFunc(value, func(c rune) bool { return c == ' ' || c == ',' })

	for _, word := range words {
		number, err := strconv.Atoi(word)
		if err == nil {
			numbers = append(numbers, number)
		}
	}

	return numbers
}

// CountTypes count number of numbers, string and others
func CountTypes(values ...interface{}) (nbNumber, nbString, nbUnknown int) {
	nbNumber, nbString, nbUnknown = 0, 0, 0

	for _, val := range values {
		switch val.(type) {
		case int:
			nbNumber++
		case float32:
			nbNumber++
		case float64:
			nbNumber++
		case string:
			nbString++
		default:
			nbUnknown++
		}
	}
	return nbNumber, nbString, nbUnknown
}

//CreateSet create a set from a list of values
func CreateSet(values ...interface{}) map[interface{}]struct{} {

	set := make(map[interface{}]struct{})

	for _, val := range values {
		set[val] = struct{}{}
	}
	return set
}

//GetEndList return the end of the list
func GetEndList(list []string, from int) ([]string, error) {
	return []string{}, nil
}
