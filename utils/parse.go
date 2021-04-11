package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func IntToString(number int) (number_string string) {
	return strconv.Itoa(number)
}

func StringToInt(number string) (number_int int) {
	if i, err := strconv.Atoi(number); err != nil {
		return -1
	} else {
		return i
	}
}

func StructToString(a interface{}) {
	out, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
