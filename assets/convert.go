package assets

import (
	"bytes"
	"strconv"
)

func ConvertArrayStringToInt(arr_str []string) []int64 {
	var arr_int = []int64{}
	for _, i := range arr_str {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		arr_int = append(arr_int, int64(j))
	}
	return arr_int
}

func ConvertNumberToWord(number int64) string {
	var word string
	arrayWord := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	word = arrayWord[number]
	return word
}
func ConvertIntToInt64Start(number int64) *int64 {
	return &number
}
func ConvertFloat64ToFloat64Start(number float64) *float64 {
	return &number
}
func NumberToString(n float64, sep rune) string {

	s := strconv.FormatFloat(n, 'f', -1, 64)

	startOffset := 0
	var buff bytes.Buffer

	if n < 0 {
		startOffset = 1
		buff.WriteByte('-')
	}

	l := len(s)

	commaIndex := 3 - ((l - startOffset) % 3)

	if commaIndex == 3 {
		commaIndex = 0
	}

	for i := startOffset; i < l; i++ {

		if commaIndex == 3 {
			buff.WriteRune(sep)
			commaIndex = 0
		}
		commaIndex++

		buff.WriteByte(s[i])
	}

	return buff.String()
}
func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func ConvertArrayToString(array []int) string {
	var result string
	for i, element := range array {
		if i+1 >= len(array) {
			str := strconv.Itoa(element)
			result += str
		} else {
			str := strconv.Itoa(element)
			result += str + ","
		}

	}
	return result
}
