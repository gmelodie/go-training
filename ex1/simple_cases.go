package ex1

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
	s := value[left:right]
	return s
}

// SubstringWithErrors return substringed value but can generate errors
func SubstringWithErrors(value string, left, right int) (string, error) {

	s := value[left:right]
	return s, nil
}

// ExtractNumbersFromString return the numbers existing in the list
func ExtractNumbersFromString(value string) []int {
	return []int{}
}

// CountTypes count number of numbers, string and others
func CountTypes(values ...interface{}) (nbNumber, nbString, nbUnknown int) {
	return 0, 0, 0
}

//CreateSet create a set from a list of values
func CreateSet(values ...interface{}) map[interface{}]struct{} {
	return make(map[interface{}]struct{})
}

//GetEndList return the end of the list
func GetEndList(list []string, from int) ([]string, error) {
	return []string{}, nil
}
