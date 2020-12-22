package Utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func ConvertIntToString(number int) (number_string string) {
	return strconv.Itoa(number)
}

func StructToString(a interface{}) {
	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
